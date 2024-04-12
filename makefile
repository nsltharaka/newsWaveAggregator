build:
	@go build -o bin/news-aggregator cmd/main.go

run: build
	@./bin/news-aggregator