package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Header struct {
	app.Compo
}

func (h *Header) Render() app.UI {
	data := []string{
		"hello",
		"go-app",
		"is",
		"sexy",
	}

	return app.H1().Body(
		app.Ul().Body(
			app.Range(data).Slice(func(i int) app.UI {
				return app.Li().Text(data[i])
			}),
		).Class("nav-items"),
	).Class("nav")
}
