package auth

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type TokenManager struct {
	secretKey []byte
	issuer    string
	blacklist *TokenBlacklist
}

func NewTokenManager(secret, issuer string) *TokenManager {
	tokenManager := &TokenManager{
		secretKey: []byte(secret),
		issuer:    issuer,
		blacklist: NewTokenBlacklist(),
	}
	
	// Start cleanup routine for expired blacklisted tokens
	tokenManager.blacklist.StartCleanupRoutine(1 * time.Hour)
	
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
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}

// InvalidateToken adds a token to the blacklist
func (tm *TokenManager) InvalidateToken(tokenString string) error {
	claims, err := tm.ValidateToken(tokenString)
	if err != nil {
		return err
	}
	
	// Add token to blacklist with its expiration time
	tm.blacklist.BlacklistToken(claims.TokenID, claims.ExpiresAt.Time)
	return nil
}