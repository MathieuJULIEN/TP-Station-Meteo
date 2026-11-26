package api

import (
	"StationMeteo/models"
	"fmt"
	"net/url"
)

var baseURL = "https://geo.api.gouv.fr"

func GetCommune(name string) error {
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

func GetVille(code string) (models.Commune, error) {
	var ville models.Commune
	err := get(baseURL, fmt.Sprintf("/commune/%v?fields=code,nom,population,centre", code), &ville)
	return ville, err
}
