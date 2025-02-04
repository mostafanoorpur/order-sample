run:
	go run cmd/server/main.go -c=config.yaml

migrate:
	goose -dir ./migrations postgres "host=localhost user=postgres dbname=postgres sslmode=disable password=postgres" up