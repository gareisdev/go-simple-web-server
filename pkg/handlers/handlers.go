package handlers

import (
	"net/http"

	"github.com/gareisdev/go-simple-web-server/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "index.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}

func Products(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "products.html")
}

func Store(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "store.html")
}
