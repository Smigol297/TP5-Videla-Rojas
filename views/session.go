package views

import "strconv"

type Resultado struct {
	Pregunta          string
	RespuestaUsuario  string
	RespuestaCorrecta string
	EsCorrecta        bool
}

func getCardClass(index int) string {
	if index == 0 {
		return "tarjeta visible"
	}
	return "tarjeta"
}
func getProgressWidth(correctas, total int) string {
	porcentaje := calcularPorcentaje(correctas, total)
	// Asegurarnos de que el porcentaje esté entre 0 y 100
	if porcentaje < 0 {
		porcentaje = 0
	}
	if porcentaje > 100 {
		porcentaje = 100
	}
	return "width: " + strconv.Itoa(porcentaje) + "%;"
}

func calcularPorcentaje(correctas, total int) int {
	if total == 0 {
		return 0
	}
	return (correctas * 100) / total
}

// getClaseResultado devuelve la clase CSS para el contenedor del resultado
func getClaseResultado(esCorrecta bool) string {
	if esCorrecta {
		return "resultado correcta"
	}
	return "resultado incorrecta"
}

// getClaseTexto devuelve la clase CSS para el texto de estado
func getClaseTexto(esCorrecta bool) string {
	if esCorrecta {
		return "correcta-text"
	}
	return "incorrecta-text"
}

// getTextoEstado devuelve el texto de estado (Correcta/Incorrecta)
func getTextoEstado(esCorrecta bool) string {
	if esCorrecta {
		return "✓ Correcta"
	}
	return "✗ Incorrecta"
}
