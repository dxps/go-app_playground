//go:build js && wasm

package ui

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func InitPWAClientSide(apiPort int) {

	initAppRoutesClientSide()
	app.RunWhenOnBrowser()
}
