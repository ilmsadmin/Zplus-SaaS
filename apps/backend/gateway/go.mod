module github.com/ilmsadmin/Zplus-SaaS/apps/backend/gateway

go 1.23.0

toolchain go1.23.9

require (
	github.com/99designs/gqlgen v0.17.74
	github.com/gofiber/fiber/v2 v2.52.0
	github.com/google/uuid v1.6.0
	github.com/ilmsadmin/Zplus-SaaS/apps/backend/shared v0.0.0-20250615015858-c6da0318b4c9
	github.com/ilmsadmin/Zplus-SaaS/pkg v0.0.0
	github.com/valyala/fasthttp v1.51.0
	github.com/vektah/gqlparser/v2 v2.5.27
	gorm.io/gorm v1.30.0
)

replace github.com/ilmsadmin/Zplus-SaaS/pkg => ../../../pkg

replace github.com/ilmsadmin/Zplus-SaaS/apps/backend/shared => ../shared

require (
	github.com/agnivade/levenshtein v1.2.1 // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.2 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/sync v0.14.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	gorm.io/driver/postgres v1.6.0 // indirect
)
