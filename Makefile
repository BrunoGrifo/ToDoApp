build:
	$ go build -o ./bin/todo cmd/main.go
run: build
	./bin/todo
migrate:
	@if [ -z "$(name)" ]; then \
		echo "Error: Migration name must be provided."; \
		echo "Usage: make migrate name=<migration_name>"; \
		exit 1; \
	fi; \
	echo "Running migration command with name: $(name)"; \
	migrate create -ext sql -dir cmd/migrate/migrations -seq $(name)

# This line ensures `make` doesn't complain about targets not being found
.PHONY: migrate

migrate_up:
	go run cmd/migrate/main.go up
migrate_down:
	go run cmd/migrate/main.go down

migrate_force:
	migrate -path cmd/migrate/migrations -database "mysql://root:<pass>@tcp(localhost:3306)/todoapp?checkConnLiveness=false&parseTime=true&maxAllowedPacket=0" force 1
