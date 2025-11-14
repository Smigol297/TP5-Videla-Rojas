BINARY=myapp

# Compilar
build:
	go build -o $(BINARY) .
# Levantar servicios con Docker Compose
up:
	@docker compose up -d
# Generar código con sqlc
sqlc:
	sqlc generate -f sqlc.yaml
# Actualizar dependencias
tidy: go.mod go.sum
	@go mod tidy
run: up tidy sqlc templ build
	./$(BINARY) & echo $$! > server.pid
	sleep 1
stop:
	@if [ -f server.pid ]; then \
		kill $$(cat server.pid); \
		rm -f server.pid; \
	else \
		 "No se encontró server.pid, utilizar lsoft -i :8080 para buscar el proceso y luego kill 'PID' para matarlo"; \
	fi

reboot: stop run

templ:
	find . -name '*_templ.go' -delete
	go run github.com/a-h/templ/cmd/templ@latest generate