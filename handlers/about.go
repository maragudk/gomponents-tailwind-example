package handlers

import (
	"net/http"

	"github.com/go-chi/chi"

	"gomponents-tailwind-example/views"
)

func About(mux chi.Router) {
	mux.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		_ = views.Page(views.Props{
			Title:   "About",
			Path:    "/about",
			Content: views.About(),
		}).Render(w)
	})
}
