package main

import (
	"context"
	"log/slog"
	"os"
	"path"
	"time"

	"github.com/dxps/go-app_playground/go-app_loading/internal/common/config"
	"github.com/dxps/go-app_playground/go-app_loading/internal/svc/api"
	"github.com/dxps/go-app_playground/go-app_loading/internal/svc/repos"
	"github.com/dxps/go-app_playground/go-app_loading/internal/svc/run"
	"github.com/dxps/go-app_playground/go-app_loading/internal/svc/uiserver"
	"github.com/sethvargo/go-envconfig"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				s := a.Value.Any().(*slog.Source)
				s.File = path.Base(s.File)
			}
			return a
		},
	}))
	slog.SetDefault(logger)

	slog.Info("Starting up ...")

	var cfg config.Config
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		slog.Error("Failed to load config.", "error", err)
		return
	}
	slog.Debug("Config loaded.")

	//////////////////////////////////////
	// API & PWA servers init & startup //
	//////////////////////////////////////

	rr := repos.NewRepos()

	apiSrv := api.NewApiServer(cfg.Servers.BackendPort)
	apiSrv.Start()
	slog.Info("Web API Server started.", "port", apiSrv.Port)

	uiSrv, err := uiserver.InitAndStartWebUiServerSide(cfg.Servers.FrontendPort, cfg.Servers.BackendPort, rr.FileRepo)
	if err != nil {
		slog.Error("Failed to start Web UI Server.", "error", err)
		return
	}
	slog.Info("Web UI Server started.", "addr", uiSrv.Addr)

	///////////////////////
	// graceful shutdown //
	///////////////////////

	shutdownCtx, shutdownCancel := context.WithCancel(context.Background())
	defer shutdownCancel()

	done := run.NewOsSignalNotifier(shutdownCtx)
	<-done.Done()
	slog.Info("Shutting down ...")

	// Give outstanding requests a deadline for completion on both API and UI servers.

	apiSrvCtx, apiSrvCancel := context.WithTimeout(shutdownCtx, 3*time.Second)
	defer apiSrvCancel()

	if err := apiSrv.Stop(apiSrvCtx); err != nil {
		slog.Error("Failed to gracefully shutdown the Web API Server.", "error", err)
	} else {
		slog.Info("Web API Server gracefully shutted down.")
	}

	uiSrvCtx, uiSrvCancel := context.WithTimeout(shutdownCtx, 3*time.Second)
	defer uiSrvCancel()

	if err := uiSrv.Shutdown(uiSrvCtx); err != nil {
		slog.Error("Failed to gracefully shutdown the Web UI Server.", "error", err)
	} else {
		slog.Info("Web UI Server gracefully shutted down.")
	}
}
