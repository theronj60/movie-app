package views

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/theronj60/movies-app/views/components"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type Layout struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *Layout) Render() app.UI {
	return app.Div().Body(
		&components.Header{},
		app.H1().
			Class("Title").
			Text("main layout file!"),
		app.P().
			Class("text").
			Text("paragraph text goes here!"),
	)
}
