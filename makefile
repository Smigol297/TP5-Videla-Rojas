BINARY=myapp

# Compilar
build:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc@latest generate
	go run github.com/a-h/templ/cmd/templ@latest generate
	go build -o $(BINARY) .

# Ejecutar servidor
run: build
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