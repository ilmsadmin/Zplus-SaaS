package auth

import (
	"sync"
	"time"
)

// TokenBlacklist manages invalidated tokens
type TokenBlacklist struct {
	blacklisted map[string]time.Time
	mutex       sync.RWMutex
}

// NewTokenBlacklist creates a new token blacklist
func NewTokenBlacklist() *TokenBlacklist {
	return &TokenBlacklist{
		blacklisted: make(map[string]time.Time),
	}
}

// BlacklistToken adds a token to the blacklist with its expiration time
func (tb *TokenBlacklist) BlacklistToken(tokenID string, expiresAt time.Time) {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	tb.blacklisted[tokenID] = expiresAt
}

// IsBlacklisted checks if a token is blacklisted
func (tb *TokenBlacklist) IsBlacklisted(tokenID string) bool {
	tb.mutex.RLock()
	defer tb.mutex.RUnlock()
	_, exists := tb.blacklisted[tokenID]
	return exists
}

// CleanupExpired removes expired tokens from the blacklist
func (tb *TokenBlacklist) CleanupExpired() {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	
	now := time.Now()
	for tokenID, expiresAt := range tb.blacklisted {
		if now.After(expiresAt) {
			delete(tb.blacklisted, tokenID)
		}
	}
}

// StartCleanupRoutine starts a goroutine that periodically cleans up expired tokens
func (tb *TokenBlacklist) StartCleanupRoutine(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		
		for range ticker.C {
			tb.CleanupExpired()
		}
	}()
}