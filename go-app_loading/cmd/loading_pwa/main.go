//go:build js && wasm

package main

import (
	"context"
	"log/slog"
	"os"
	"path"
	"strings"

	"github.com/dxps/go-app_playground/go-app_loading/internal/common/config"
	"github.com/dxps/go-app_playground/go-app_loading/internal/ui"
	"github.com/sethvargo/go-envconfig"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				s := a.Value.Any().(*slog.Source)
				// Log only the filename, without the extension.
				// s.File = path.Base(s.File)
				s.File, _ = strings.CutSuffix(path.Base(s.File), ".go")
			} else
			// Remove the timestamp.
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			} else
			// Remove the logging level.
			if a.Key == slog.LevelKey {
				return slog.Attr{}
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

	/////////////////////
	// PWA server init //
	/////////////////////

	ui.InitPWAClientSide(cfg.Servers.BackendPort)
}
