package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/dxps/go-app_playground/go-app_loading/internal/common/model"
)

func (s *ApiServer) getHealthCheck(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	data := model.Health{State: "ok"}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("Failed to encode response as json.", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
