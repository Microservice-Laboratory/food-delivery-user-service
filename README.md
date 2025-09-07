# food-delivery-user-service

## How to run test

1. go clean -testcache
2. go test ./... -coverprofile=coverage.out -coverpkg=./...
3. go tool cover -html=coverage.out

## Handle migrations

1. To create: migrate create -ext sql -dir db/migrations -seq migration_name
2. To run: migrate -database ${POSTGRESQL_URL} -path db/migrations up

## Run application

1. go run cmd/main.go
