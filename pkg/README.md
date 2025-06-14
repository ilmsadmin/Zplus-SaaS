# Shared Packages

This directory contains reusable Go packages and JavaScript/TypeScript libraries.

## Go Packages

### Auth Package (`auth/`)
- JWT token management
- Token generation and validation
- Claims handling for multi-tenant authentication

### Database Package (`database/`)
- Database connection utilities
- Tenant schema switching
- Connection pooling helpers

### Utils Package (`utils/`)
- Common utility functions
- String manipulation helpers
- Validation functions

## Usage

Import packages in your Go services:

```go
import (
    "github.com/ilmsadmin/Zplus-SaaS/pkg/auth"
    "github.com/ilmsadmin/Zplus-SaaS/pkg/database" 
    "github.com/ilmsadmin/Zplus-SaaS/pkg/utils"
)

// Create token manager
tm := auth.NewTokenManager("secret", "zplus-saas")

// Generate JWT token
token, err := tm.GenerateToken(userID, tenantID, role)

// Connect to database
db, err := database.Connect(config)

// Set tenant context
err = database.SetTenantSchema(db, tenantID)
```

## JavaScript/TypeScript SDKs

Future JavaScript/TypeScript SDKs for frontend applications will be added here:
- API client libraries
- Type definitions
- Utility functions for frontend apps

## Development

This package uses Go modules. To add dependencies:

```bash
cd pkg
go mod tidy
```