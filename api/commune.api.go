package api

import (
	"StationMeteo/models"
	"fmt"
	"net/url"
)

func GetCommunes(name string) ([]models.Commune, error) {
	var baseURL = "https://geo.api.gouv.fr"
	params := url.Values{}
	if name != "" {
		params.Set("name", name)
	}
	params.Set("fields", "code,nom,population,centre")
	endpoint := "/communes?" + params.Encode()
	var communes []models.Commune
	err := get(baseURL, endpoint, &communes)
	return communes, err
}

func GetCommuneByCode(code string) (models.Commune, error) {
	var baseURL = "https://geo.api.gouv.fr"
	var commune models.Commune
	err := get(baseURL, fmt.Sprintf("/communes/%v?fields=code,nom,population,centre", code), &commune)
	return commune, err
}
