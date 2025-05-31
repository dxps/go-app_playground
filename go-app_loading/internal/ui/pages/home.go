package pages

import (
	"github.com/dxps/go-app_playground/go-app_loading/internal/ui/comps"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type HomePage struct {
	app.Compo
}

func (p *HomePage) OnMount(ctx app.Context) {

	qps := ctx.Page().URL().Query()
	if qps.Has("backto") {
		ctx.Navigate(qps.Get("backto"))
	}
}

func (h *HomePage) Render() app.UI {
	return app.Div().Class(comps.PageCss).Body(
		&comps.Navbar{},
		app.Div().
			Class(comps.MainContentCss).
			Body(
				app.Div().Class(comps.MainContentChildDivCss).Body(
					app.Div().Class("flex flex-col items-center my-8").Body(
						app.Div().Body(app.Raw(comps.LogoIcon)).Class("w-12"),
						app.Div().Text("Loading sample").Class("text-3xl text-gray-400"),
						app.Div().Text("A sample that showcases how to customize the loading screen.").Class("my-4 text-gray-700"),
					),
				),
			),
	)
}
