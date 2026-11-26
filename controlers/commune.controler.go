package controllers

import (
	"StationMeteo/helpers"
	"StationMeteo/models"
	"StationMeteo/services"
	"errors"
	"log"
	"net/http"
)

func (tabtemp *models.Centre) MeteoCommune(w http.ResponseWriter, r *http.Request) {
	tabtempe := tabtemp.Coordinate
	commune, err := services.GetMeteo(tabtempe[1], tabtempe[0])
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
