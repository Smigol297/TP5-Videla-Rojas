package logic

import (
	"fmt"
	sqlc "tp5/db"

	_ "github.com/lib/pq"
)

// INICIO CON MINUSCULA = NO PUBLICO
// INICIO CON MAYUSCULA = PUBLICO
func ValidateCreateTarjeta(p sqlc.CreateTarjetaParams) error {
	if p.Pregunta == "" {
		return fmt.Errorf("la pregunta no puede estar vacía")
	}
	if p.Respuesta == "" {
		return fmt.Errorf("la respuesta no puede estar vacía")
	}
	if p.OpcionA == "" {
		return fmt.Errorf("la opción A no puede estar vacía")
	}
	if p.OpcionB == "" {
		return fmt.Errorf("la opción B no puede estar vacía")
	}
	if p.OpcionC == "" {
		return fmt.Errorf("la opción C no puede estar vacía")
	}
	if p.IDTema <= 0 {
		return fmt.Errorf("ID de tema inválido")
	}
	return nil
}
func ValidateUpdateTarjeta(p sqlc.UpdateTarjetaParams) error {
	if p.IDTarjeta <= 0 {
		return fmt.Errorf("ID de tarjeta %d inválido", p.IDTarjeta)
	}
	if p.Pregunta == "" {
		return fmt.Errorf("la pregunta no puede estar vacía")
	}
	if p.Respuesta == "" {
		return fmt.Errorf("la respuesta no puede estar vacía")
	}
	if p.OpcionA == "" {
		return fmt.Errorf("la opción A no puede estar vacía")
	}
	if p.OpcionB == "" {
		return fmt.Errorf("la opción B no puede estar vacía")
	}
	if p.OpcionC == "" {
		return fmt.Errorf("la opción C no puede estar vacía")
	}
	if p.IDTema <= 0 {
		return fmt.Errorf("ID de tema inválido")
	}
	return nil
}
