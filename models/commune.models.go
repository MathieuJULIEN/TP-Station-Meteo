package models

type Commune struct {
	Code       string  `json:"code"`
	Nom        string  `json:"nom"`
	Population int     `json:"population"`
	Centre     *Centre `json:"centre"`
}

type Centre struct {
	Type       string    `json:"type"`
	Coordinate []float64 `json:"coordinates"`
}

func (c Centre) Longitude() float64 {
	var lat float64
	lat = c.Coordinate[0]
	return lat
}

func (c Centre) Latitude() float64 {
	var long float64
	long = c.Coordinate[1]
	return long
}
