package domain

type Buses struct {
	IdBus        int `json:"idBus"`
	Placa       string `json:"placa"`
	Capacidad   int `json:"capacidad"`
	ChoferID    int `json:"choferID"`
}

func NewBus(idBus int, placa string, capacidad int, choferID int) *Buses {
	return &Buses{
		IdBus:      idBus, 
		Placa:      placa,
		Capacidad:  capacidad,
		ChoferID:   choferID, 
	}
}