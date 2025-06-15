package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/shared/models"
)

// UserService handles CRUD operations for tenant users
type UserService struct {
	db       *gorm.DB
	tenantID uuid.UUID
}

// NewUserService creates a new user service for a specific tenant
func NewUserService(db *gorm.DB, tenantID uuid.UUID) *UserService {
	return &UserService{
		db:       db,
		tenantID: tenantID,
	}
}

// CreateUserInput represents input for creating a user
type CreateUserInput struct {
	Email     string      `json:"email" validate:"required,email"`
	Password  string      `json:"password" validate:"required,min=8"`
	FirstName string      `json:"first_name" validate:"required"`
	LastName  string      `json:"last_name" validate:"required"`
	Avatar    *string     `json:"avatar"`
	RoleIDs   []uuid.UUID `json:"role_ids"`
}

// UpdateUserInput represents input for updating a user
type UpdateUserInput struct {
	Email     *string `json:"email"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Avatar    *string `json:"avatar"`
	Status    *string `json:"status"`
}

// UserFilter represents filtering options for users
type UserFilter struct {
	Status string `json:"status"`
	Role   string `json:"role"`
	Search string `json:"search"`
}

// CreateUser creates a new user within the tenant
func (s *UserService) CreateUser(input CreateUserInput) (*models.TenantUser, error) {
	// Validate email is unique within tenant
	var existingCount int64
	if err := s.db.Model(&models.TenantUser{}).
		Where("tenant_id = ? AND email = ?", s.tenantID, input.Email).
		Count(&existingCount).Error; err != nil {
		return nil, fmt.Errorf("failed to check email uniqueness: %v", err)
	}

	if existingCount > 0 {
		return nil, fmt.Errorf("user with email '%s' already exists in this tenant", input.Email)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	user := &models.TenantUser{
		TenantID:     s.tenantID,
		Email:        strings.ToLower(strings.TrimSpace(input.Email)),
		PasswordHash: string(hashedPassword),
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Avatar:       input.Avatar,
		Status:       "active",
	}

	// Create user
	if err := s.db.Create(user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	// Assign roles if provided
	if len(input.RoleIDs) > 0 {
		if err := s.assignRolesToUser(user.ID, input.RoleIDs); err != nil {
			return nil, fmt.Errorf("failed to assign roles: %v", err)
		}
	}

	// Load user with roles
	if err := s.db.Preload("Roles").Preload("Roles.Permissions").First(user, user.ID).Error; err != nil {
		return nil, fmt.Errorf("failed to load user details: %v", err)
	}

	return user, nil
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id uuid.UUID) (*models.TenantUser, error) {
	var user models.TenantUser
	err := s.db.Preload("Roles").Preload("Roles.Permissions").
		Where("tenant_id = ?", s.tenantID).
		First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	return &user, nil
}

// GetUserByEmail retrieves a user by email
func (s *UserService) GetUserByEmail(email string) (*models.TenantUser, error) {
	var user models.TenantUser
	err := s.db.Preload("Roles").Preload("Roles.Permissions").
		Where("tenant_id = ? AND email = ?", s.tenantID, strings.ToLower(email)).
		First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	return &user, nil
}

// ListUsers retrieves users with filtering and pagination
func (s *UserService) ListUsers(filter UserFilter, offset, limit int) ([]*models.TenantUser, int64, error) {
	query := s.db.Model(&models.TenantUser{}).
		Preload("Roles").
		Where("tenant_id = ?", s.tenantID)

	// Apply filters
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.Search != "" {
		search := "%" + filter.Search + "%"
		query = query.Where("first_name ILIKE ? OR last_name ILIKE ? OR email ILIKE ?", search, search, search)
	}
	if filter.Role != "" {
		query = query.Joins("JOIN user_roles ON users.id = user_roles.user_id").
			Joins("JOIN roles ON user_roles.role_id = roles.id").
			Where("roles.name = ?", filter.Role)
	}

	// Get total count
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count users: %v", err)
	}

	// Get paginated results
	var users []*models.TenantUser
	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&users).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list users: %v", err)
	}

	return users, total, nil
}

// UpdateUser updates a user
func (s *UserService) UpdateUser(id uuid.UUID, input UpdateUserInput) (*models.TenantUser, error) {
	var user models.TenantUser
	if err := s.db.Where("tenant_id = ?", s.tenantID).First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to find user: %v", err)
	}

	// Update fields
	if input.Email != nil {
		email := strings.ToLower(strings.TrimSpace(*input.Email))
		// Check if email is unique within tenant (excluding current user)
		var existingCount int64
		if err := s.db.Model(&models.TenantUser{}).
			Where("tenant_id = ? AND email = ? AND id != ?", s.tenantID, email, id).
			Count(&existingCount).Error; err != nil {
			return nil, fmt.Errorf("failed to check email uniqueness: %v", err)
		}
		if existingCount > 0 {
			return nil, fmt.Errorf("user with email '%s' already exists in this tenant", email)
		}
		user.Email = email
	}
	if input.FirstName != nil {
		user.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		user.LastName = *input.LastName
	}
	if input.Avatar != nil {
		user.Avatar = input.Avatar
	}
	if input.Status != nil {
		user.Status = *input.Status
	}

	if err := s.db.Save(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	// Load user with roles
	if err := s.db.Preload("Roles").Preload("Roles.Permissions").First(&user, user.ID).Error; err != nil {
		return nil, fmt.Errorf("failed to load user details: %v", err)
	}

	return &user, nil
}

// DeleteUser soft deletes a user
func (s *UserService) DeleteUser(id uuid.UUID) error {
	var user models.TenantUser
	if err := s.db.Where("tenant_id = ?", s.tenantID).First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("failed to find user: %v", err)
	}

	// Soft delete the user
	if err := s.db.Delete(&user).Error; err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	return nil
}

// ChangePassword changes a user's password
func (s *UserService) ChangePassword(id uuid.UUID, oldPassword, newPassword string) error {
	var user models.TenantUser
	if err := s.db.Where("tenant_id = ?", s.tenantID).First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("failed to find user: %v", err)
	}

	// Verify old password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(oldPassword)); err != nil {
		return fmt.Errorf("current password is incorrect")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash new password: %v", err)
	}

	// Update password
	if err := s.db.Model(&user).Update("password_hash", string(hashedPassword)).Error; err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}

	return nil
}

// AssignRoles assigns roles to a user
func (s *UserService) AssignRoles(userID uuid.UUID, roleIDs []uuid.UUID) error {
	// Verify user exists and belongs to tenant
	var user models.TenantUser
	if err := s.db.Where("tenant_id = ?", s.tenantID).First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("failed to find user: %v", err)
	}

	return s.assignRolesToUser(userID, roleIDs)
}

// RemoveRoles removes roles from a user
func (s *UserService) RemoveRoles(userID uuid.UUID, roleIDs []uuid.UUID) error {
	// Verify user exists and belongs to tenant
	var user models.TenantUser
	if err := s.db.Where("tenant_id = ?", s.tenantID).First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("failed to find user: %v", err)
	}

	// Remove role assignments
	if err := s.db.Where("user_id = ? AND role_id IN ?", userID, roleIDs).
		Delete(&models.UserRole{}).Error; err != nil {
		return fmt.Errorf("failed to remove roles: %v", err)
	}

	return nil
}

// UpdateLastLogin updates the user's last login timestamp
func (s *UserService) UpdateLastLogin(id uuid.UUID) error {
	now := time.Now()
	result := s.db.Model(&models.TenantUser{}).
		Where("tenant_id = ? AND id = ?", s.tenantID, id).
		Update("last_login_at", now)
	
	if result.Error != nil {
		return fmt.Errorf("failed to update last login: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	
	return nil
}

// Helper methods

func (s *UserService) assignRolesToUser(userID uuid.UUID, roleIDs []uuid.UUID) error {
	// Verify roles exist and belong to tenant
	var validRoles []models.Role
	if err := s.db.Where("tenant_id = ? AND id IN ?", s.tenantID, roleIDs).
		Find(&validRoles).Error; err != nil {
		return fmt.Errorf("failed to verify roles: %v", err)
	}

	if len(validRoles) != len(roleIDs) {
		return fmt.Errorf("some roles were not found or don't belong to this tenant")
	}

	// Remove existing role assignments
	if err := s.db.Where("user_id = ?", userID).Delete(&models.UserRole{}).Error; err != nil {
		return fmt.Errorf("failed to remove existing roles: %v", err)
	}

	// Create new role assignments
	now := time.Now()
	for _, roleID := range roleIDs {
		userRole := models.UserRole{
			UserID:     userID,
			RoleID:     roleID,
			AssignedAt: now,
		}
		if err := s.db.Create(&userRole).Error; err != nil {
			return fmt.Errorf("failed to assign role: %v", err)
		}
	}

	return nil
}