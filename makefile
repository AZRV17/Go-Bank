run:
	go run ./cmd

migrateup:
	migrate -path pkg/db/migrations -database "postgresql://postgres:sa@localhost:5432/gowebdb?sslmode=disable" -verbose up

migratedown:
	migrate -path pkg/db/migrations -database "postgresql://postgres:sa@localhost:5432/gowebdb?sslmode=disable" -verbose down

test:
	go test ./... -v -cover
