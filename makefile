run:
	go run ./cmd

migrateup:
	migrate -path pkg/db/migrations -database "postgresql://postgres:sa@localhost:5432/gowebdb?sslmode=disable" -verbose up

migratedown:
	migrate -path pkg/db/migrations -database "postgresql://postgres:sa@localhost:5432/gowebdb?sslmode=disable" -verbose down

test:
	go test ./... -v -cover

coverage:
	go test ./... -v -coverprofile=coverage.out

proto:
	protoc -I pkg/proto pkg/proto/* --go-grpc_out=internal/server/ --go_out=internal/server/
