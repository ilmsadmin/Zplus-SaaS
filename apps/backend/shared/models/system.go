package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SystemUser represents a system-level administrator
type SystemUser struct {
	ID           uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email        string         `json:"email" gorm:"unique;not null"`
	PasswordHash string         `json:"-" gorm:"not null"`
	Name         string         `json:"name" gorm:"not null"`
	Role         string         `json:"role" gorm:"not null"` // super_admin, admin, support
	IsActive     bool           `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName returns the table name for SystemUser
func (SystemUser) TableName() string {
	return "system.system_users"
}

// Tenant represents a tenant/organization in the system
type Tenant struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name       string         `json:"name" gorm:"not null"`
	Slug       string         `json:"slug" gorm:"unique;not null"`
	Domain     *string        `json:"domain"`
	Subdomain  *string        `json:"subdomain"`
	PlanID     *uuid.UUID     `json:"plan_id" gorm:"type:uuid"`
	Plan       *Plan          `json:"plan,omitempty" gorm:"foreignKey:PlanID"`
	Status     string         `json:"status" gorm:"default:'active'"` // active, suspended, trial, expired
	Settings   map[string]interface{} `json:"settings" gorm:"type:jsonb;default:'{}'"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName returns the table name for Tenant
func (Tenant) TableName() string {
	return "system.tenants"
}

// Plan represents a subscription plan
type Plan struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string         `json:"name" gorm:"not null"`
	Description *string        `json:"description"`
	Price       float64        `json:"price" gorm:"type:decimal(10,2)"`
	Features    map[string]interface{} `json:"features" gorm:"type:jsonb;default:'{}'"`
	MaxUsers    *int           `json:"max_users"`
	MaxStorage  *int64         `json:"max_storage"` // bytes
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName returns the table name for Plan
func (Plan) TableName() string {
	return "system.plans"
}

// Module represents available system modules
type Module struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string         `json:"name" gorm:"not null"`
	Description *string        `json:"description"`
	Version     string         `json:"version" gorm:"default:'1.0.0'"`
	Enabled     bool           `json:"enabled" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName returns the table name for Module
func (Module) TableName() string {
	return "system.modules"
}

// TenantModule represents which modules are enabled for a tenant
type TenantModule struct {
	TenantID      uuid.UUID `json:"tenant_id" gorm:"type:uuid;primaryKey"`
	ModuleID      uuid.UUID `json:"module_id" gorm:"type:uuid;primaryKey"`
	Tenant        Tenant    `json:"tenant,omitempty" gorm:"foreignKey:TenantID"`
	Module        Module    `json:"module,omitempty" gorm:"foreignKey:ModuleID"`
	Enabled       bool      `json:"enabled" gorm:"default:true"`
	Configuration map[string]interface{} `json:"configuration" gorm:"type:jsonb;default:'{}'"`
	CreatedAt     time.Time `json:"created_at"`
}

// TableName returns the table name for TenantModule
func (TenantModule) TableName() string {
	return "system.tenant_modules"  
}

// Subscription represents a tenant's subscription to a plan
type Subscription struct {
	ID            uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TenantID      uuid.UUID      `json:"tenant_id" gorm:"type:uuid;not null"`
	PlanID        uuid.UUID      `json:"plan_id" gorm:"type:uuid;not null"`
	Tenant        Tenant         `json:"tenant,omitempty" gorm:"foreignKey:TenantID"`
	Plan          Plan           `json:"plan,omitempty" gorm:"foreignKey:PlanID"`
	Status        string         `json:"status" gorm:"default:'active'"` // active, cancelled, expired, trial
	StartDate     time.Time      `json:"start_date"`
	EndDate       *time.Time     `json:"end_date"`
	TrialEndDate  *time.Time     `json:"trial_end_date"`
	BillingCycle  string         `json:"billing_cycle" gorm:"default:'monthly'"` // monthly, yearly
	Amount        float64        `json:"amount" gorm:"type:decimal(10,2)"`
	Currency      string         `json:"currency" gorm:"default:'USD'"`
	Metadata      map[string]interface{} `json:"metadata" gorm:"type:jsonb;default:'{}'"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName returns the table name for Subscription
func (Subscription) TableName() string {
	return "system.subscriptions"
}