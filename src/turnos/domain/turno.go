package domain

import "time"

type Turno struct {
	ID       int `json:"id"`
	ChoferID int `json:"choferID"`
	BusID    int `json:"busID"`
	Inicio   time.Time `json:"inicio"`
	Fin      *time.Time `json:"fin"`
	Estado   string `json:"estado"`
}

