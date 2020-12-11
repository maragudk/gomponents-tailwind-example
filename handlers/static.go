package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Static(mux chi.Router) {
	// Serve static files from the "public" directory
	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("public")))
	mux.Get("/static/*", staticHandler.ServeHTTP)
}
