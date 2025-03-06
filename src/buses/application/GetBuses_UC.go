package application

import "bus-driver/src/buses/domain"

type GetBusesUC struct {
	db domain.IBusesRepository
}

func NewGetBusesUC(db domain.IBusesRepository) *GetBusesUC {
	return &GetBusesUC{db: db}
}

func (uc *GetBusesUC) Run() ([]domain.Buses, error) {
	return uc.db.GetBuses()
}