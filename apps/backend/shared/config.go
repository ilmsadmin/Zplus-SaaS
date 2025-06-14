package shared

// Database configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

// JWT configuration
type JWTConfig struct {
	Secret     string
	ExpiresIn  int
	RefreshIn  int
}

// Redis configuration
type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

// Application configuration
type Config struct {
	Database DatabaseConfig
	JWT      JWTConfig
	Redis    RedisConfig
}