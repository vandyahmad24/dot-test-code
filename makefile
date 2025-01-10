run-api:
	go run cmd/api/main.go

migrate-up:
	go run cmd/migration/main.go up

migrate-down:
	go run cmd/migration/main.go down

migrate-status:
	go run cmd/migration/main.go status

migrate-create:
	@echo "Creating migration file for $(migration_name)"
	go run cmd/migration/main.go create $(migration_name)
