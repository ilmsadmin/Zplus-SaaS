package auth

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type TokenManager struct {
	secretKey      []byte
	issuer         string
	blacklist      *TokenBlacklist
	sessionManager *SessionManager
}

func NewTokenManager(secret, issuer string) *TokenManager {
	tokenManager := &TokenManager{
		secretKey:      []byte(secret),
		issuer:         issuer,
		blacklist:      NewTokenBlacklist(),
		sessionManager: NewSessionManager(),
	}
	
	// Start cleanup routines
	tokenManager.blacklist.StartCleanupRoutine(1 * time.Hour)
	tokenManager.sessionManager.StartCleanupRoutine(1 * time.Hour, 24 * time.Hour) // Clean sessions idle for 24 hours
	
	return tokenManager
}

type Claims struct {
	UserID   string `json:"user_id"`
	TenantID string `json:"tenant_id"`
	Role     string `json:"role"`
	TokenID  string `json:"token_id"` // Add unique token ID for blacklisting
	jwt.RegisteredClaims
}

func (tm *TokenManager) GenerateToken(userID, tenantID, role string) (string, error) {
	tokenID := uuid.New().String()
	expiresAt := time.Now().Add(24 * time.Hour)
	
	claims := Claims{
		UserID:   userID,
		TenantID: tenantID,
		Role:     role,
		TokenID:  tokenID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    tm.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(tm.secretKey)
}

func (tm *TokenManager) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return tm.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// Check if token is blacklisted
		if tm.blacklist.IsBlacklisted(claims.TokenID) {
			return nil, jwt.ErrTokenInvalidClaims
		}
		
		// Update session activity
		tm.sessionManager.UpdateLastSeen(claims.TokenID)
		
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}

// InvalidateToken adds a token to the blacklist and removes associated session
func (tm *TokenManager) InvalidateToken(tokenString string) error {
	claims, err := tm.ValidateToken(tokenString)
	if err != nil {
		return err
	}
	
	// Add token to blacklist with its expiration time
	tm.blacklist.BlacklistToken(claims.TokenID, claims.ExpiresAt.Time)
	
	// Remove associated session
	tm.sessionManager.RemoveSession(claims.TokenID)
	
	return nil
}

// CreateSession creates a new session for a token
func (tm *TokenManager) CreateSession(tokenID, userID, tenantID, email, ipAddress, userAgent string) *Session {
	return tm.sessionManager.CreateSession(tokenID, userID, tenantID, email, ipAddress, userAgent)
}

// GetSession retrieves a session by token ID
func (tm *TokenManager) GetSession(tokenID string) *Session {
	return tm.sessionManager.GetSession(tokenID)
}

// UpdateSessionActivity updates the last seen time for a session
func (tm *TokenManager) UpdateSessionActivity(tokenID string) {
	tm.sessionManager.UpdateLastSeen(tokenID)
}

// GetUserSessions retrieves all active sessions for a user
func (tm *TokenManager) GetUserSessions(userID string) []*Session {
	return tm.sessionManager.GetUserSessions(userID)
}

// GetAllSessions retrieves all active sessions
func (tm *TokenManager) GetAllSessions() []*Session {
	return tm.sessionManager.GetAllSessions()
}