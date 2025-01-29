module airline-tracking-service

go 1.22

toolchain go1.23.5

require (
	github.com/joho/godotenv v1.5.1
	github.com/redis/go-redis/v9 v9.7.0
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
)

replace golang.org/x/sys => golang.org/x/sys v0.7.0
