package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"tp5/db"
	"tp5/views"
)

func VerificarRespuestasHandler(w http.ResponseWriter, r *http.Request, queries *db.Queries, temaId int) {
	err := r.ParseForm() //Le dice a Go: "Lee todos los datos que vinieron en la petición POST y organízalos".
	if err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	var resultados []views.Resultado
	correctas := 0
	total := 0

	// Procesar cada respuesta
	for key, values := range r.Form {
		if strings.HasPrefix(key, "respuesta") && !strings.HasPrefix(key, "respuestaCorrecta") {
			idxStr := strings.TrimPrefix(key, "respuesta")
			_, err := strconv.Atoi(idxStr)
			if err != nil {
				continue
			}

			respuestaUsuario := strings.TrimSpace(strings.ToUpper(values[0]))
			respuestaCorrecta := strings.TrimSpace(strings.ToUpper(r.Form.Get("respuestaCorrecta" + idxStr)))
			pregunta := r.Form.Get("pregunta" + idxStr)

			esCorrecta := respuestaUsuario == respuestaCorrecta
			if esCorrecta {
				correctas++
			}
			total++

			resultados = append(resultados, views.Resultado{
				Pregunta:          pregunta,
				RespuestaUsuario:  values[0],
				RespuestaCorrecta: respuestaCorrecta,
				EsCorrecta:        esCorrecta,
			})
		}
	}

	// Renderizar template Templ
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.ResultsPage(resultados, correctas, total, temaId).Render(r.Context(), w)
}
