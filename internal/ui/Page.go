package ui

import "github.com/rivo/tview"

type Page interface {
	GetPrimitive() tview.Primitive
	OnFocus()                // Optional: Aktionen, wenn Seite Fokus bekommt
	OnBlur()                 // Optional: Aktionen, wenn Seite Fokus verliert
	RenderWithData(data any) // Optional: Seite mit Daten rendern
}
