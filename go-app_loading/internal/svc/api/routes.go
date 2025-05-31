package api

import apiroutes "github.com/dxps/go-app_playground/go-app_loading/internal/common/api_routes"

// initRoutes initializes the routes that the API Server will serve.
func (s *ApiServer) initRoutes() {
	s.router.Get(apiroutes.Health, s.getHealthCheck)
}
