package services

import (
	"StationMeteo/api"
	"StationMeteo/models"
)

func GetMeteo(lat float64, long float64) (models.Meteo, error) {
	return api.GetForecast(lat, long)
}
