# food-delivery-user-service

## How to run test

1. go clean -testcache
2. go test ./... -coverprofile=coverage.out -coverpkg=./...
3. go tool cover -html=coverage.out
