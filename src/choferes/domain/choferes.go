package domain

type Chofer struct {
	ID        int `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido_p string `json:"apellido_p"`
	Apellido_m string `json:"apellido_m"`
	Edad	  int `json:"edad"`
}

func NewChofer(id int, nombre string, apellido_p string, apellido_m string, edad int) *Chofer {
	return &Chofer{
		ID:       id,
		Nombre:   nombre,
		Apellido_p: apellido_p,
		Apellido_m: apellido_m,
		Edad:     edad,
	}
}