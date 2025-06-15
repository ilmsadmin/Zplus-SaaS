package auth

import (
	"sync"
	"time"
)

// Session represents an active user session
type Session struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	TenantID  string    `json:"tenant_id"`
	TokenID   string    `json:"token_id"`
	Email     string    `json:"email"`
	IPAddress string    `json:"ip_address"`
	UserAgent string    `json:"user_agent"`
	CreatedAt time.Time `json:"created_at"`
	LastSeen  time.Time `json:"last_seen"`
}

// SessionManager manages active user sessions
type SessionManager struct {
	sessions map[string]*Session // key: tokenID
	userSessions map[string][]string // key: userID, value: slice of tokenIDs
	mutex    sync.RWMutex
}

// NewSessionManager creates a new session manager
func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions:     make(map[string]*Session),
		userSessions: make(map[string][]string),
	}
}

// CreateSession creates a new session for a user
func (sm *SessionManager) CreateSession(tokenID, userID, tenantID, email, ipAddress, userAgent string) *Session {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	
	session := &Session{
		ID:        tokenID, // Use tokenID as session ID for simplicity
		UserID:    userID,
		TenantID:  tenantID,
		TokenID:   tokenID,
		Email:     email,
		IPAddress: ipAddress,
		UserAgent: userAgent,
		CreatedAt: time.Now(),
		LastSeen:  time.Now(),
	}
	
	sm.sessions[tokenID] = session
	
	// Add to user sessions
	if sm.userSessions[userID] == nil {
		sm.userSessions[userID] = make([]string, 0)
	}
	sm.userSessions[userID] = append(sm.userSessions[userID], tokenID)
	
	return session
}

// GetSession retrieves a session by token ID
func (sm *SessionManager) GetSession(tokenID string) *Session {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	
	session, exists := sm.sessions[tokenID]
	if !exists {
		return nil
	}
	
	return session
}

// UpdateLastSeen updates the last seen time for a session
func (sm *SessionManager) UpdateLastSeen(tokenID string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	
	if session, exists := sm.sessions[tokenID]; exists {
		session.LastSeen = time.Now()
	}
}

// RemoveSession removes a session
func (sm *SessionManager) RemoveSession(tokenID string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	
	session, exists := sm.sessions[tokenID]
	if !exists {
		return
	}
	
	// Remove from sessions
	delete(sm.sessions, tokenID)
	
	// Remove from user sessions
	userSessions := sm.userSessions[session.UserID]
	for i, sessionID := range userSessions {
		if sessionID == tokenID {
			sm.userSessions[session.UserID] = append(userSessions[:i], userSessions[i+1:]...)
			break
		}
	}
	
	// Clean up empty user session slice
	if len(sm.userSessions[session.UserID]) == 0 {
		delete(sm.userSessions, session.UserID)
	}
}

// GetUserSessions retrieves all active sessions for a user
func (sm *SessionManager) GetUserSessions(userID string) []*Session {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	
	tokenIDs, exists := sm.userSessions[userID]
	if !exists {
		return nil
	}
	
	sessions := make([]*Session, 0, len(tokenIDs))
	for _, tokenID := range tokenIDs {
		if session, exists := sm.sessions[tokenID]; exists {
			sessions = append(sessions, session)
		}
	}
	
	return sessions
}

// GetAllSessions retrieves all active sessions
func (sm *SessionManager) GetAllSessions() []*Session {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	
	sessions := make([]*Session, 0, len(sm.sessions))
	for _, session := range sm.sessions {
		sessions = append(sessions, session)
	}
	
	return sessions
}

// CleanupExpiredSessions removes sessions that haven't been seen for a specified duration
func (sm *SessionManager) CleanupExpiredSessions(maxIdleTime time.Duration) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	
	cutoff := time.Now().Add(-maxIdleTime)
	expiredTokens := make([]string, 0)
	
	for tokenID, session := range sm.sessions {
		if session.LastSeen.Before(cutoff) {
			expiredTokens = append(expiredTokens, tokenID)
		}
	}
	
	// Remove expired sessions
	for _, tokenID := range expiredTokens {
		session := sm.sessions[tokenID]
		delete(sm.sessions, tokenID)
		
		// Remove from user sessions
		userSessions := sm.userSessions[session.UserID]
		for i, sessionID := range userSessions {
			if sessionID == tokenID {
				sm.userSessions[session.UserID] = append(userSessions[:i], userSessions[i+1:]...)
				break
			}
		}
		
		// Clean up empty user session slice
		if len(sm.userSessions[session.UserID]) == 0 {
			delete(sm.userSessions, session.UserID)
		}
	}
}

// StartCleanupRoutine starts a goroutine that periodically cleans up expired sessions
func (sm *SessionManager) StartCleanupRoutine(interval time.Duration, maxIdleTime time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		
		for range ticker.C {
			sm.CleanupExpiredSessions(maxIdleTime)
		}
	}()
}