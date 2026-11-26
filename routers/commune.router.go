package routes

import (
	"StationMeteo/controlers"
	"net/http"
)

func communeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /commune", controllers.ListCommunes)
	mux.HandleFunc("GET /commune/{code}", controllers.MeteoCommune)
}
