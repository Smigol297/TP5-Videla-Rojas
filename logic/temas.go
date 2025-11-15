package logic

import (
	"fmt"
	sqlc "tp5/db"

	_ "github.com/lib/pq"
)

// INICIO CON MINUSCULA = NO PUBLICO
// INICIO CON MAYUSCULA = PUBLICO
func ValidateCreateTema(nombre string) error {
	if nombre == "" {
		return fmt.Errorf("el nombre del tema no puede estar vacío")
	}
	return nil
}
func ValidateUpdateTema(p sqlc.UpdateTemaParams) error {
	if p.IDTema <= 0 {
		return fmt.Errorf("ID de tema %d inválido", p.IDTema)
	}
	if p.NombreTema == "" {
		return fmt.Errorf("el nombre del tema no puede estar vacío")
	}

	return nil
}
