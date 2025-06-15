package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"
)

type LoginRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	TenantSlug string `json:"tenant_slug"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	User         interface{} `json:"user"`
	ExpiresIn    int    `json:"expires_in"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func TestFullAuthenticationFlow(t *testing.T) {
	// Skip if running in CI without services
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping integration tests")
	}

	// Start auth service in background
	authCmd := exec.Command("go", "run", "main.go")
	authCmd.Dir = "../auth"
	authCmd.Env = append(os.Environ(), "PORT=8081")
	
	if err := authCmd.Start(); err != nil {
		t.Fatalf("Failed to start auth service: %v", err)
	}
	defer authCmd.Process.Kill()

	// Wait for auth service to start
	time.Sleep(2 * time.Second)

	// Test 1: Login with valid credentials
	loginReq := LoginRequest{
		Email:      "admin@demo-corp.zplus.com",
		Password:   "demo123",
		TenantSlug: "demo-corp",
	}

	loginBody, _ := json.Marshal(loginReq)
	resp, err := http.Post("http://localhost:8081/login", "application/json", bytes.NewBuffer(loginBody))
	if err != nil {
		t.Fatalf("Login request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("Expected login status 200, got %d. Response: %s", resp.StatusCode, string(body))
	}

	var loginResp LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		t.Fatalf("Failed to parse login response: %v", err)
	}

	if loginResp.Token == "" {
		t.Fatal("Expected token in login response")
	}

	fmt.Printf("✓ Login successful, token: %s...\n", loginResp.Token[:20])

	// Test 2: Logout with valid token
	logoutReq, _ := http.NewRequest("POST", "http://localhost:8081/logout", nil)
	logoutReq.Header.Set("Authorization", "Bearer "+loginResp.Token)

	client := &http.Client{}
	logoutResp, err := client.Do(logoutReq)
	if err != nil {
		t.Fatalf("Logout request failed: %v", err)
	}
	defer logoutResp.Body.Close()

	if logoutResp.StatusCode != 200 {
		body, _ := io.ReadAll(logoutResp.Body)
		t.Fatalf("Expected logout status 200, got %d. Response: %s", logoutResp.StatusCode, string(body))
	}

	fmt.Println("✓ Logout successful")

	// Test 3: Try to logout again with the same token (should fail)
	logoutReq2, _ := http.NewRequest("POST", "http://localhost:8081/logout", nil)
	logoutReq2.Header.Set("Authorization", "Bearer "+loginResp.Token)

	invalidResp, err := client.Do(logoutReq2)
	if err != nil {
		t.Fatalf("Second logout request failed: %v", err)
	}
	defer invalidResp.Body.Close()

	if invalidResp.StatusCode != 401 {
		body, _ := io.ReadAll(invalidResp.Body)
		t.Fatalf("Expected status 401 for invalidated token, got %d. Response: %s", invalidResp.StatusCode, string(body))
	}

	fmt.Println("✓ Token invalidation working correctly")

	// Test 4: Login with invalid credentials
	invalidLoginReq := LoginRequest{
		Email:      "nonexistent@demo-corp.zplus.com",
		Password:   "wrongpassword",
		TenantSlug: "demo-corp",
	}

	invalidLoginBody, _ := json.Marshal(invalidLoginReq)
	invalidLoginResp, err := http.Post("http://localhost:8081/login", "application/json", bytes.NewBuffer(invalidLoginBody))
	if err != nil {
		t.Fatalf("Invalid login request failed: %v", err)
	}
	defer invalidLoginResp.Body.Close()

	if invalidLoginResp.StatusCode != 401 {
		t.Fatalf("Expected status 401 for invalid credentials, got %d", invalidLoginResp.StatusCode)
	}

	fmt.Println("✓ Invalid credentials properly rejected")

	// Test 5: Logout without token
	logoutReq3, _ := http.NewRequest("POST", "http://localhost:8081/logout", nil)
	noTokenResp, err := client.Do(logoutReq3)
	if err != nil {
		t.Fatalf("Logout without token request failed: %v", err)
	}
	defer noTokenResp.Body.Close()

	if noTokenResp.StatusCode != 400 {
		t.Fatalf("Expected status 400 for missing token, got %d", noTokenResp.StatusCode)
	}

	fmt.Println("✓ Logout without token properly rejected")
}

func TestRefreshTokenFlow(t *testing.T) {
	// Skip if running in CI without services
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping integration tests")
	}

	// Start auth service in background
	authCmd := exec.Command("go", "run", "main.go")
	authCmd.Dir = "../auth"
	authCmd.Env = append(os.Environ(), "PORT=8081")
	
	if err := authCmd.Start(); err != nil {
		t.Fatalf("Failed to start auth service: %v", err)
	}
	defer authCmd.Process.Kill()

	// Wait for auth service to start
	time.Sleep(2 * time.Second)

	// Login first
	loginReq := LoginRequest{
		Email:      "admin@demo-corp.zplus.com",
		Password:   "demo123",
		TenantSlug: "demo-corp",
	}

	loginBody, _ := json.Marshal(loginReq)
	resp, err := http.Post("http://localhost:8081/login", "application/json", bytes.NewBuffer(loginBody))
	if err != nil {
		t.Fatalf("Login request failed: %v", err)
	}
	defer resp.Body.Close()

	var loginResp LoginResponse
	json.NewDecoder(resp.Body).Decode(&loginResp)

	// Test refresh token
	refreshReq := map[string]string{
		"refresh_token": loginResp.RefreshToken,
	}

	refreshBody, _ := json.Marshal(refreshReq)
	refreshResp, err := http.Post("http://localhost:8081/refresh", "application/json", bytes.NewBuffer(refreshBody))
	if err != nil {
		t.Fatalf("Refresh request failed: %v", err)
	}
	defer refreshResp.Body.Close()

	if refreshResp.StatusCode != 200 {
		body, _ := io.ReadAll(refreshResp.Body)
		t.Fatalf("Expected refresh status 200, got %d. Response: %s", refreshResp.StatusCode, string(body))
	}

	var refreshRespData LoginResponse
	if err := json.NewDecoder(refreshResp.Body).Decode(&refreshRespData); err != nil {
		t.Fatalf("Failed to parse refresh response: %v", err)
	}

	if refreshRespData.Token == "" {
		t.Fatal("Expected new token in refresh response")
	}

	fmt.Println("✓ Token refresh working correctly")
}