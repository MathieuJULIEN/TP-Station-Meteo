package models

type Meteo struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Daily Daily `json:"daily"`
}

type Daily struct {
	Time []string `json:"time"`
	Temperature []int `json:"temperature_2m_mean"`
}