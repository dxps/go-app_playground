package uiserver

import (
	"github.com/dxps/go-app_playground/go-app_loading/internal/common"
	"github.com/dxps/go-app_playground/go-app_loading/internal/ui/pages"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func initAppHomeRoutesServerSide() {
	app.Route(common.HomePath, func() app.Composer { return &pages.HomePage{} })
}
