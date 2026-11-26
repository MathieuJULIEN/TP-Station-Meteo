package routes

import (
	"StationMeteo/controlers"
	"net/http"
)

func characterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /commune", )
	mux.HandleFunc("GET /commune/{code}", controllers.MeteoCommune)
}
