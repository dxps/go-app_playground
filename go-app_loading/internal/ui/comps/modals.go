package comps

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

const (
	ConfirmModalCss = `fixed inset-0 flex flex-wrap justify-center items-center w-full h-full z-[1000] 
	                   before:fixed before:inset-0 before:w-full before:h-full before:bg-[rgba(0,0,0,0.5)] overflow-auto`
)

////////////////////////
// confirmation modal //
////////////////////////

type ConfirmationModal struct {
	app.Compo
	title           string
	content         string
	confirmBtnLabel string
	confirmCallback func(ctx app.Context, e app.Event)
	show            *bool
}

func NewConfirmationModal(
	title string,
	content string,
	confirmBtnLabel string,
	confirmCallback func(ctx app.Context, e app.Event),
	show *bool,
) *ConfirmationModal {
	return &ConfirmationModal{
		title:           title,
		content:         content,
		confirmBtnLabel: confirmBtnLabel,
		confirmCallback: confirmCallback,
		show:            show,
	}
}

func (c *ConfirmationModal) Render() app.UI {

	return app.Div().Class(ConfirmModalCss).Body(
		app.Div().Class(MainContentCss).Body(
			app.Div().Class(MainContentChildDivCss).Body(
				app.Div().Class("flex").Body(
					app.P().Text(c.title).Class(TitleCss),
					app.P().Text("x").
						Class(CloseSymbolCss).
						OnClick(func(ctx app.Context, e app.Event) { *c.show = false }),
				),
				app.Div().Class("mt-8 text-gray-600 flex justify-center").Body(
					app.P().Text(c.content),
				),
				app.Div().Class("grid justify-items-end mt-8").Body(
					app.Button().Text(c.confirmBtnLabel).Class(DeleteBtnConfirmCss).
						OnClick(c.confirmCallback),
				),
			),
		),
	)
}

////////////////
// info modal //
////////////////

type InfoModal struct {
	app.Compo
	title   string
	content string
	show    *bool
}

func NewInfoModal(title string, content string, show *bool) *InfoModal {
	return &InfoModal{
		title:   title,
		content: content,
		show:    show,
	}
}

func (i *InfoModal) Render() app.UI {
	if !*i.show {
		return app.Div()
	}
	return app.Div().Class(ConfirmModalCss).Body(
		app.Div().Class(MainContentCss).Body(
			app.Div().Class(MainContentChildDivCss).Body(
				app.Div().Class("flex").Body(
					app.P().Text(i.title).Class(TitleCss),
					app.P().Text("x").
						Class(CloseSymbolCss).
						OnClick(func(ctx app.Context, e app.Event) { i.Hide() }),
				),
				app.Div().Class("mt-8 text-gray-600 flex justify-center").Body(
					app.P().Text(i.content),
				),
				app.Div().Class("grid justify-items-end mt-8").Body(
					app.Button().Text("OK").Class(CloseBtnCss).
						OnClick(func(ctx app.Context, e app.Event) { i.Hide() }),
				),
			),
		),
	)
}

func (i *InfoModal) Show() {
	*i.show = true
}

func (i *InfoModal) Hide() {
	*i.show = false
}
