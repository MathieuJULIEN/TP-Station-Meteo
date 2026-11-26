package routes

import (
	controllers "StationMeteo/controlers"
	"net/http"
)

func MainRouter() *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	mux.Handle("GET /{$}", http.RedirectHandler("/commune", http.StatusMovedPermanently))

	communeRoutes(mux)

	mux.HandleFunc("/", controllers.NotFound)
	return mux
}
