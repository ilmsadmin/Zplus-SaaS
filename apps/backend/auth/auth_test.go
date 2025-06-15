package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/auth/handlers"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/auth/models"
)

func TestLoginLogoutFlow(t *testing.T) {
	// Create a new Fiber app for testing
	app := fiber.New()
	authHandler := handlers.NewAuthHandler()

	// Register routes
	app.Post("/login", authHandler.Login)
	app.Post("/logout", authHandler.Logout)

	// Test login
	loginReq := models.LoginRequest{
		Email:      "admin@demo-corp.zplus.com",
		Password:   "demo123",
		TenantSlug: "demo-corp",
	}

	loginBody, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, 5000)
	if err != nil {
		t.Fatalf("Login request failed: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}

	// Parse login response
	var loginResp models.LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		t.Fatalf("Failed to parse login response: %v", err)
	}

	if loginResp.Token == "" {
		t.Fatal("Expected token in login response")
	}

	fmt.Printf("✓ Login successful, token: %s...\n", loginResp.Token[:20])

	// Test logout with token
	logoutReq, _ := http.NewRequest("POST", "/logout", nil)
	logoutReq.Header.Set("Authorization", "Bearer "+loginResp.Token)

	logoutResp, err := app.Test(logoutReq, 5000)
	if err != nil {
		t.Fatalf("Logout request failed: %v", err)
	}

	if logoutResp.StatusCode != 200 {
		t.Fatalf("Expected logout status 200, got %d", logoutResp.StatusCode)
	}

	fmt.Println("✓ Logout successful")

	// Test using the token again (should fail)
	loginReq2, _ := http.NewRequest("POST", "/logout", nil)
	loginReq2.Header.Set("Authorization", "Bearer "+loginResp.Token)

	invalidResp, err := app.Test(loginReq2, 5000)
	if err != nil {
		t.Fatalf("Second logout request failed: %v", err)
	}

	if invalidResp.StatusCode != 401 {
		t.Fatalf("Expected status 401 for invalidated token, got %d", invalidResp.StatusCode)
	}

	fmt.Println("✓ Token invalidation working correctly")
}

func TestLoginWithInvalidCredentials(t *testing.T) {
	app := fiber.New()
	authHandler := handlers.NewAuthHandler()
	app.Post("/login", authHandler.Login)

	loginReq := models.LoginRequest{
		Email:      "nonexistent@demo-corp.zplus.com",
		Password:   "wrongpassword",
		TenantSlug: "demo-corp",
	}

	loginBody, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, 5000)
	if err != nil {
		t.Fatalf("Login request failed: %v", err)
	}

	if resp.StatusCode != 401 {
		t.Fatalf("Expected status 401 for invalid credentials, got %d", resp.StatusCode)
	}

	fmt.Println("✓ Invalid credentials properly rejected")
}

func TestLogoutWithoutToken(t *testing.T) {
	app := fiber.New()
	authHandler := handlers.NewAuthHandler()
	app.Post("/logout", authHandler.Logout)

	req, _ := http.NewRequest("POST", "/logout", nil)
	resp, err := app.Test(req, 5000)
	if err != nil {
		t.Fatalf("Logout request failed: %v", err)
	}

	if resp.StatusCode != 400 {
		t.Fatalf("Expected status 400 for missing token, got %d", resp.StatusCode)
	}

	fmt.Println("✓ Logout without token properly rejected")
}

func TestSessionManagement(t *testing.T) {
	app := fiber.New()
	authHandler := handlers.NewAuthHandler()

	app.Post("/login", authHandler.Login)
	app.Get("/sessions", authHandler.GetSessions)
	app.Post("/logout", authHandler.Logout)

	// Test login to create session
	loginReq := models.LoginRequest{
		Email:      "admin@demo-corp.zplus.com",
		Password:   "demo123",
		TenantSlug: "demo-corp",
	}

	loginBody, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, 5000)
	if err != nil {
		t.Fatalf("Login request failed: %v", err)
	}

	var loginResp models.LoginResponse
	json.NewDecoder(resp.Body).Decode(&loginResp)

	// Check sessions endpoint
	sessionsReq, _ := http.NewRequest("GET", "/sessions", nil)
	sessionsResp, err := app.Test(sessionsReq, 5000)
	if err != nil {
		t.Fatalf("Sessions request failed: %v", err)
	}

	if sessionsResp.StatusCode != 200 {
		t.Fatalf("Expected sessions status 200, got %d", sessionsResp.StatusCode)
	}

	var sessionsRespData map[string]interface{}
	json.NewDecoder(sessionsResp.Body).Decode(&sessionsRespData)

	count := int(sessionsRespData["count"].(float64))
	if count != 1 {
		t.Fatalf("Expected 1 active session, got %d", count)
	}

	fmt.Println("✓ Session created successfully")

	// Test logout to remove session
	logoutReq, _ := http.NewRequest("POST", "/logout", nil)
	logoutReq.Header.Set("Authorization", "Bearer "+loginResp.Token)

	app.Test(logoutReq, 5000)

	// Check sessions again
	sessionsReq2, _ := http.NewRequest("GET", "/sessions", nil)
	sessionsResp2, err := app.Test(sessionsReq2, 5000)
	if err != nil {
		t.Fatalf("Sessions request failed: %v", err)
	}

	var sessionsRespData2 map[string]interface{}
	json.NewDecoder(sessionsResp2.Body).Decode(&sessionsRespData2)

	count2 := int(sessionsRespData2["count"].(float64))
	if count2 != 0 {
		t.Fatalf("Expected 0 active sessions after logout, got %d", count2)
	}

	fmt.Println("✓ Session removed successfully after logout")
}