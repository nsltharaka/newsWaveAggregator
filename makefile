build:
	@go build -o bin/news-aggregator cmd/main.go

run: build
	@./bin/news-aggregator

goose:
	@goose -dir sql/schema postgres "postgres://postgres:root@localhost:5432/aggregatordb?sslmode=disable" status

goose-up:
	@goose -dir sql/schema postgres "postgres://postgres:root@localhost:5432/aggregatordb?sslmode=disable" up

goose-down:
	@goose -dir sql/schema postgres "postgres://postgres:root@localhost:5432/aggregatordb?sslmode=disable" down

goose-reset:
	@goose -dir sql/schema postgres "postgres://postgres:root@localhost:5432/aggregatordb?sslmode=disable" reset

goose-update: goose-reset goose-up

goose-data:
	@goose -dir sql/data postgres "postgres://postgres:root@localhost:5432/aggregatordb?sslmode=disable" up

sc :
	@sqlc generate
