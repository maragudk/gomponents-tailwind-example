package server

import (
	"gomponents-tailwind-example/handlers"
)

// setupRoutes and middleware.
func (s *Server) setupRoutes() {
	handlers.Static(s.mux)
	handlers.Home(s.mux)
	handlers.About(s.mux)
}
