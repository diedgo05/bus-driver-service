package domain

type IBusesRepository interface {
	Save(bus Buses) error
	GetBuses() ([]Buses, error)
	FindBusByIdChofer(choferID int) ([]Buses, error)
	UpdateByID(idBus int, bus Buses) error
	DeleteByID(idBus int) error
}