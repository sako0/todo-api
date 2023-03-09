migrate-up:
	go run cmd/migration/main.go

migrate-down:
	go run cmd/migration/main.go down

# テストを実施
test:
	go test -v ./...
