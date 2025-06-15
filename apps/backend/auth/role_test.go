package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/auth/handlers"
	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/auth/models"
)

func setupRoleTestApp() *fiber.App {
	app := fiber.New()
	roleHandler := handlers.NewRoleHandler()

	// Role management endpoints
	app.Get("/roles", roleHandler.GetRoles)
	app.Get("/roles/:id", roleHandler.GetRole)
	app.Post("/roles", roleHandler.CreateRole)
	app.Put("/roles/:id", roleHandler.UpdateRole)
	app.Delete("/roles/:id", roleHandler.DeleteRole)

	// Permission management endpoints
	app.Get("/permissions", roleHandler.GetPermissions)
	app.Post("/permissions", roleHandler.CreatePermission)

	// Role-Permission assignment endpoints
	app.Get("/roles/:id/permissions", roleHandler.GetRolePermissions)
	app.Post("/roles/permissions", roleHandler.AssignPermissionToRole)

	// User-Role assignment endpoints
	app.Get("/users/:id/roles", roleHandler.GetUserRoles)
	app.Post("/users/roles", roleHandler.AssignRoleToUser)

	return app
}

func TestGetRoles(t *testing.T) {
	app := setupRoleTestApp()

	req, _ := http.NewRequest("GET", "/roles", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	if response["count"] == nil {
		t.Error("Expected count field in response")
	}

	if response["data"] == nil {
		t.Error("Expected data field in response")
	}

	// Should have default roles
	count := int(response["count"].(float64))
	if count < 4 {
		t.Errorf("Expected at least 4 default roles, got %d", count)
	}
}

func TestCreateRole(t *testing.T) {
	app := setupRoleTestApp()

	roleData := models.CreateRoleRequest{
		Name:         "test_role",
		DisplayName:  "Test Role",
		Description:  "A test role",
		IsSystemRole: false,
	}

	jsonData, _ := json.Marshal(roleData)
	req, _ := http.NewRequest("POST", "/roles", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode != fiber.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	if response["data"] == nil {
		t.Error("Expected data field in response")
	}

	role := response["data"].(map[string]interface{})
	if role["name"] != "test_role" {
		t.Errorf("Expected role name 'test_role', got %v", role["name"])
	}
}

func TestCreateDuplicateRole(t *testing.T) {
	app := setupRoleTestApp()

	// Try to create a role with existing name
	roleData := models.CreateRoleRequest{
		Name:         "system_admin", // This already exists
		DisplayName:  "Another System Admin",
		Description:  "Duplicate role",
		IsSystemRole: false,
	}

	jsonData, _ := json.Marshal(roleData)
	req, _ := http.NewRequest("POST", "/roles", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode != fiber.StatusConflict {
		t.Errorf("Expected status 409, got %d", resp.StatusCode)
	}
}

func TestGetRole(t *testing.T) {
	app := setupRoleTestApp()

	// Get role with ID 1 (system_admin)
	req, _ := http.NewRequest("GET", "/roles/1", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	if response["data"] == nil {
		t.Error("Expected data field in response")
	}

	role := response["data"].(map[string]interface{})
	if role["name"] != "system_admin" {
		t.Errorf("Expected role name 'system_admin', got %v", role["name"])
	}
}

func TestGetNonExistentRole(t *testing.T) {
	app := setupRoleTestApp()

	req, _ := http.NewRequest("GET", "/roles/999", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode != fiber.StatusNotFound {
		t.Errorf("Expected status 404, got %d", resp.StatusCode)
	}
}

func TestUpdateRole(t *testing.T) {
	app := setupRoleTestApp()

	updateData := models.UpdateRoleRequest{
		DisplayName: "Updated Role Name",
		Description: "Updated description",
	}

	jsonData, _ := json.Marshal(updateData)
	req, _ := http.NewRequest("PUT", "/roles/2", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	role := response["data"].(map[string]interface{})
	if role["display_name"] != "Updated Role Name" {
		t.Errorf("Expected updated display name, got %v", role["display_name"])
	}
}

func TestDeleteSystemRole(t *testing.T) {
	app := setupRoleTestApp()

	// Try to delete system_admin role (ID 1)
	req, _ := http.NewRequest("DELETE", "/roles/1", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode != fiber.StatusForbidden {
		t.Errorf("Expected status 403, got %d", resp.StatusCode)
	}
}

func TestGetPermissions(t *testing.T) {
	app := setupRoleTestApp()

	req, _ := http.NewRequest("GET", "/permissions", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	if response["count"] == nil {
		t.Error("Expected count field in response")
	}

	// Should have default permissions
	count := int(response["count"].(float64))
	if count < 10 {
		t.Errorf("Expected at least 10 default permissions, got %d", count)
	}
}

func TestCreatePermission(t *testing.T) {
	app := setupRoleTestApp()

	permData := models.CreatePermissionRequest{
		Name:        "reports:generate",
		Resource:    "reports",
		Action:      "generate",
		Description: "Generate reports",
	}

	jsonData, _ := json.Marshal(permData)
	req, _ := http.NewRequest("POST", "/permissions", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode != fiber.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	perm := response["data"].(map[string]interface{})
	if perm["name"] != "reports:generate" {
		t.Errorf("Expected permission name 'reports:generate', got %v", perm["name"])
	}
}

func TestAssignRoleToUser(t *testing.T) {
	app := setupRoleTestApp()

	assignData := models.AssignRoleRequest{
		UserID: "test-user-1",
		RoleID: 1,
	}

	jsonData, _ := json.Marshal(assignData)
	req, _ := http.NewRequest("POST", "/users/roles", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestAssignPermissionToRole(t *testing.T) {
	app := setupRoleTestApp()

	assignData := models.AssignPermissionRequest{
		RoleID:       1,
		PermissionID: 1,
	}

	jsonData, _ := json.Marshal(assignData)
	req, _ := http.NewRequest("POST", "/roles/permissions", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	// This should fail with conflict since permission is already assigned
	if resp.StatusCode != fiber.StatusConflict {
		t.Errorf("Expected status 409, got %d", resp.StatusCode)
	}
}

func TestGetRolePermissions(t *testing.T) {
	app := setupRoleTestApp()

	// Get permissions for system_admin role (ID 1)
	req, _ := http.NewRequest("GET", "/roles/1/permissions", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	if response["permissions"] == nil {
		t.Error("Expected permissions field in response")
	}

	permissions := response["permissions"].([]interface{})
	if len(permissions) == 0 {
		t.Error("Expected system_admin to have permissions")
	}
}

func TestGetUserRoles(t *testing.T) {
	app := setupRoleTestApp()

	// First assign a role to a user
	assignData := models.AssignRoleRequest{
		UserID: "test-user-2",
		RoleID: 2,
	}

	jsonData, _ := json.Marshal(assignData)
	req, _ := http.NewRequest("POST", "/users/roles", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	app.Test(req)

	// Now get user roles
	req, _ = http.NewRequest("GET", "/users/test-user-2/roles", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	if response["roles"] == nil {
		t.Error("Expected roles field in response")
	}

	roles := response["roles"].([]interface{})
	if len(roles) != 1 {
		t.Errorf("Expected 1 role, got %d", len(roles))
	}
}