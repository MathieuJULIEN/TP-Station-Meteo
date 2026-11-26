package api

import (
	"StationMeteo/models"
	"fmt"
	"net/url"
)

func GetCommunes(name string) error {
	var baseURL = "https://geo.api.gouv.fr"
	endpoint := "/commune"
	if name != "" {
		params := url.Values{}
		params.Set("name", name)
		endpoint += "?" + params.Encode()
	}
	var communes models.Commune
	err := get(baseURL, endpoint, &communes)
	return err
}

func GetCommuneByCode(code string) (models.Commune, error) {
	var baseURL = "https://geo.api.gouv.fr"
	var commune models.Commune
	err := get(baseURL, fmt.Sprintf("/commune/%v?fields=code,nom,population,centre", code), &commune)
	return commune, err
}
