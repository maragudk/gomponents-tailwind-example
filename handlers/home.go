package handlers

import (
	"net/http"

	"github.com/go-chi/chi"

	"gomponents-tailwind-example/views"
)

func Home(mux chi.Router) {
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// This is just to show how to map dynamic data to view components
		routes := mux.Routes()

		_ = views.Page(views.Props{
			Title:   "Home",
			Path:    "/",
			Content: views.Home(routes),
		}).Render(w)
	})
}
