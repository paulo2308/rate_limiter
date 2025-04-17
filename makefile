# Configurações
BINARY_NAME=rate_limiter
MAIN_FILE=cmd/main.go

# Comandos

run:
	go run $(MAIN_FILE)

redis-up:
	docker-compose up -d

redis-down:
	docker-compose down

