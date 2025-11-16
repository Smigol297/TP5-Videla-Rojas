# TP3/4 - Videla Rojas

## 📌 Requisitos

### 1. Golang
### 2. Docker
### 3. SQLC

## 🚀 Instalación y ejecución del proyecto

1. Clonar el proyecto

```bash
git clone https://github.com/Smigol297/TP5-Videla-Rojas
cd TP5-Videla-Rojas
```

2. Iniciar 

```bash
make run
```


3. Detener

```bash
make stop
```

## 📝 Tests
Ejecutar todos los tests
```bash
make allTests
```
Test usuarios
```bash
make testUsers
```
Test tarjetas
```bash
make testTarjetas
```
Test temas
```bash
make testTemas
```
## 📝 Comandos individuales
### Usuarios
Listar usuarios
```bash
make listUsuarios
```
Crear usuario
```bash
make createUsuario
```
Obtener usuario por ID
```bash
make getUserByID
```
Modificar usuario por ID
```bash
make putUserByID
```
Eliminar usuario por ID
```bash
make deleteUserByID
```
### Tarjetas
Listar tarjetas
```bash
make listTarjetas
```
Listar tarjetas por tema
```bash
make listTarjetasByTema
```
Crear tarjeta
```bash
make createTarjeta
```
Obtener tarjeta por ID
```bash
make getTarjetaByID
```
Modificar tarjeta por ID
```bash
make putTarjetaByID
```
Eliminar tarjeta por ID
```bash
make deleteTarjetaByID
```
### Temas
Listar temas
```bash
make listTemas
```
Crear tema
```bash
make createTema
```
Obtener tema por ID
```bash
make getTemaByID
```
Modificar tema por ID
```bash
make putTemaByID
```
Eliminar tema por ID
```bash
make deleteTemaByID
```
## 🛑 Detener contenedores
Detener docker
```bash
make down
```

