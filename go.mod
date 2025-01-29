module airline-tracking-service

go 1.22

toolchain go1.23.5

require (
	github.com/joho/godotenv v1.5.1
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
)

replace golang.org/x/sys => golang.org/x/sys v0.7.0
