package handler

import (
	"github.com/tsurugi-dev/go-bookings/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.Render(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.Render(w, "about.page.tmpl")
}
