package views

import (
	"tp5/db"

	"github.com/a-h/templ"
)

func IndexContent(data any, temas []db.Tema) templ.Component {
	if temas != nil {
		switch d := data.(type) {
		case []db.Tarjetum:
			return TarjetaAndTemaBody(d, temas)
		default:
			return EmptyPage()
		}
	}

	switch d := data.(type) {
	case []db.Tarjetum:
		return TarjetaBody(d)
	case []db.Usuario:
		return UserBody(d)
	case db.Tarjetum:
		return TarjetaIDBody(d)
	default:
		return EmptyPage()
	}
}
