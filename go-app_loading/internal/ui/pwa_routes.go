package ui

import (
	"github.com/dxps/go-app_playground/go-app_loading/internal/ui/pages"
	"github.com/dxps/go-app_playground/go-app_loading/internal/ui/uiroutes"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func initAppRoutesClientSide() {

	app.Route(uiroutes.Home, func() app.Composer { return &pages.HomePage{} })
}
