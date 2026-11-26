package services

import (
	"StationMeteo/api"
	"StationMeteo/models"
)

func GetMeteo(lat float64, long float64) (models.Meteo, error) {
	return api.GetForecast(lat, long)
}

func SearchCommune(name string) ([]models.Commune, error) {
	return api.GetCommunes(name)
}

func GetCommune(code string) (models.Commune, error) {
	return api.GetCommuneByCode(code)
}
