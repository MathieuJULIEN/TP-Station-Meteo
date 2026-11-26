package controllers

import (
	"StationMeteo/helpers"
	"StationMeteo/models"
	"StationMeteo/services"
	"errors"
	"log"
	"net/http"
	"strings"
)

func ListCommunes(w http.ResponseWriter, r *http.Request) {
	search := strings.TrimSpace(r.URL.Query().Get("name"))
	communes, err := services.
	if err != nil {
		renderAPIError(w, err)
		return
	}
	page.Communes = communes
	helpers.Render(w, http.StatusOK, "list", ListPage{Search : search, Communes : communes})
}

func MeteoCommune(w http.ResponseWriter, r *http.Request) {
	
	meteo, err := services.GetMeteo(Commune.Centre.Latitude(), Commune.Centre.Longitude())
	if err != nil {
		renderAPIError(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "commune", commune)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	helpers.RenderError(w, http.StatusNotFound, "Cette page n'existe pas.")
}

func renderAPIError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, models.ErrNotFound):
		helpers.RenderError(w, http.StatusNotFound,
			"Cette commune n'existe pas.")
	case errors.Is(err, models.ErrBadRequest):
		helpers.RenderError(w, http.StatusBadRequest,
			"Code de commune invalide.")
	case errors.Is(err, models.ErrUnavailable):
		log.Println("API :", err)
		helpers.RenderError(w, http.StatusServiceUnavailable,
			"Service momentanément indisponible.")
	default:
		log.Println("Erreur inattendue :", err)
		helpers.RenderError(w, http.StatusInternalServerError,
			"Une erreur est survenue.")
	}
}
