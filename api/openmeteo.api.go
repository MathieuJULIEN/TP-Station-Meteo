package api

import (
	"StationMeteo/models"
	"fmt"
)

func GetForecast(latitude float64, longitude float64) (models.Meteo, error){
	var baseURL = "https://api.open-meteo.com/v1"
	var meteo models.Meteo
	err := get(baseURL, fmt.Sprintf("/forecast?latitude=%f&longitude=%f&daily=temperature_2m_mean", latitude, longitude), &meteo)
	return meteo, err
}