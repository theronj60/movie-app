package views

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Movies struct {
	app.Compo
}

func (h *Movies) Render() app.UI {
	return app.H1().Text("Movie page goes here")
	}
