module github.com/ilmsadmin/Zplus-SaaS/apps/backend/auth

go 1.21

require (
	github.com/gofiber/fiber/v2 v2.52.0
	github.com/ilmsadmin/Zplus-SaaS/pkg v0.0.0
	golang.org/x/crypto v0.31.0
)

replace github.com/ilmsadmin/Zplus-SaaS/pkg => ../../../pkg

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
)
