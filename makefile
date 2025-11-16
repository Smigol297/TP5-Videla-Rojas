BINARY=myapp

# Compilar
build:
	go build -o $(BINARY) .
# Levantar servicios con Docker Compose
up:
	@docker compose up -d
	sleep 1
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
		 "No se encontró server.pid, utilizar lsof -i :8080 para buscar el proceso y luego kill 'PID' para matarlo"; \
	fi
down:
	@docker compose down
	@rm -rf $(BINARY)

reboot: stop run

templ:
	find . -name '*_templ.go' -delete
	go run github.com/a-h/templ/cmd/templ@latest generate

#createTarjeta:
#Ejemplo: curl -X POST http://localhost:8081/tarjetas \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "pregunta=¿Capital de Francia actualizada?" \
  -d "respuesta=París" \
  -d "opcion-a=París" \
  -d "opcion-b=Berlín" \
  -d "opcion-c=Madrid" \
  -d "id-tema=1"
#updateTarjeta:
#Ejemplo: curl -X POST http://localhost:8081/tarjetas \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "pregunta=¿Capital de Francia actualizada?" \
  -d "respuesta=París" \
  -d "opcion-a=París" \
  -d "opcion-b=Berlín" \
  -d "opcion-c=Madrid" \
  -d "id-tema=1"
#deleteTarjeta:
#Ejemplo: curl -X DELETE http://localhost:8081/tarjetas/1