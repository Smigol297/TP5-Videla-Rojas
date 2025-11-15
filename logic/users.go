package logic

import (
	"fmt"
	"net/mail"
	sqlc "tp5/db"

	_ "github.com/lib/pq"
)

// INICIO CON MINUSCULA = NO PUBLICO
// INICIO CON MAYUSCULA = PUBLICO
func ValidateCreateUser(p sqlc.CreateUsuarioParams) error {
	if p.NombreUsuario == "" {
		return fmt.Errorf("el nombre del usuario no puede estar vacío")
	}
	_, err := mail.ParseAddress(p.Email)
	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("el email del usuario no es válido")
	}
	if p.Contrasena == "" {
		return fmt.Errorf("la contraseña del usuario no puede estar vacía")
	}
	return nil
}
func ValidateUpdateUser(p sqlc.UpdateUsuarioParams) error {
	if p.IDUsuario <= 0 {
		return fmt.Errorf("ID de usuario %d inválido", p.IDUsuario)
	}
	if p.NombreUsuario == "" {
		return fmt.Errorf("el nombre del usuario no puede estar vacío")
	}
	_, err := mail.ParseAddress(p.Email)
	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("el email del usuario no es válido")
	}
	if p.Contrasena == "" {
		return fmt.Errorf("la contraseña del usuario no puede estar vacía")
	}
	return nil
}
