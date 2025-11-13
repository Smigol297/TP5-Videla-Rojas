-- 1. Crear Usuario (no depende de nada)
CREATE TABLE Usuario (
    id_usuario SERIAL PRIMARY KEY,
    nombre_usuario VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    contrasena VARCHAR(255) NOT NULL
);

-- 2. Crear Tema (no depende de nada, PERO Tarjeta depende de ella)
CREATE TABLE Tema (
    id_tema SERIAL PRIMARY KEY,
    nombre_tema VARCHAR(100) NOT NULL UNIQUE
);

-- 3. Crear Tarjeta (DEBE IR AL FINAL, porque depende de Tema)
CREATE TABLE Tarjeta (
    id_tarjeta SERIAL PRIMARY KEY,
    pregunta VARCHAR(255) NOT NULL,
    respuesta VARCHAR(255) NOT NULL,
    opcion_a VARCHAR(255) NOT NULL,
    opcion_b VARCHAR(255) NOT NULL,
    opcion_c VARCHAR(255) NOT NULL,
    id_tema INTEGER NOT NULL,
    CONSTRAINT fk_tema
        FOREIGN KEY(id_tema) 
        REFERENCES Tema(id_tema)
        ON DELETE CASCADE
);
