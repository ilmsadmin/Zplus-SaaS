package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	SSLMode  string
}

func Connect(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Host, config.Username, config.Password, config.Database, config.Port, config.SSLMode)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}

func SetTenantSchema(db *gorm.DB, tenantID string) error {
	schemaName := fmt.Sprintf("tenant_%s", tenantID)
	return db.Exec(fmt.Sprintf("SET search_path TO %s", schemaName)).Error
}