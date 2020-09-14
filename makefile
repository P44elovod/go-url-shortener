run:
	go run ./cmd/main.go
postgres:
	docker run --name psql -p 5432\:5432 -e POSTGRES_USER\=root -e POSTGRES_PASSWORD\=pass -d postgres\:13-alpine
	psqlstart:
	docker start psql
psqlstop:
	docker stop psql
createdb:
	docker exec -it psql createdb --username=root --owner=root urldb
dropdb:
	docker exec -it psql dropdb urldb
migrateup:
	migrate -path migrations -database "postgresql://root:pass@localhost:5432/mts?sslmode=disable" -verbose up
migratedown:
	migrate -path migrations -database "postgresql://root:pass@localhost:5432/mts?sslmode=disable" -verbose down
.PHONY: run

.DEFAULT_GOAL := run