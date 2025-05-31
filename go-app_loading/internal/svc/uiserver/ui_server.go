package uiserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dxps/go-app_playground/go-app_loading/internal/svc/repos"
)

func InitAndStartWebUiServerSide(uiPort, apiPort int, fileRepo *repos.FileRepo) (*http.Server, error) {

	initAppHomeRoutesServerSide()

	handler, err := newCustomAppHandler(fileRepo)
	if err != nil {
		return nil, err
	}
	uiSrv := http.Server{
		Addr:    fmt.Sprintf(":%d", uiPort),
		Handler: handler,
	}

	go func() {
		if err := uiSrv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	return &uiSrv, nil
}
