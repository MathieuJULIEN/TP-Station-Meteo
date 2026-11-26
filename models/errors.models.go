package models

import "errors"

var (
	ErrNotFound    = errors.New("ressource introuvable")
	ErrBadRequest  = errors.New("demande invalide")
	ErrUnavailable = errors.New("API indisponible")
)
