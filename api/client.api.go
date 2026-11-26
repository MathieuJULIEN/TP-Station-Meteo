package api

import (
	"StationMeteo/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 5,
}

func get(baseURL string, endpoint string, destination any) error {
	request, requestErr := http.NewRequest(http.MethodGet, baseURL+endpoint, nil)
	if requestErr != nil {
		return fmt.Errorf("création de la requete : %w", requestErr)
	}
	response, responseErr := client.Do(request)
	if responseErr != nil {
		return fmt.Errorf("%w : %v", models.ErrUnavailable, responseErr)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Erreur requete (%w) : %v", models.ErrBadRequest, response.StatusCode)
	}

	err := json.NewDecoder(response.Body).Decode(destination)
	if err != nil {
		return fmt.Errorf("décodage de la réponse : %w", err)
	}
	return nil
}

