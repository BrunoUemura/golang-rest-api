# golang-rest-api

golang-rest-api

## Setup

### Dependencies

- go mod init github.com/BrunoUemura/golang-rest-api
- go get github.com/gin-gonic/gin
- go get gorm.io/gorm
- go get gorm.io/driver/postgres

### Start Application

Using docker

- docker-compose build
- docker-compose up -d

Locally (Need docker for database)

- docker-compose up -d
- go run main.go
