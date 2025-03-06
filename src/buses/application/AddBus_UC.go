package application

import "bus-driver/src/buses/domain"

type AddBusUC struct {
	db domain.IBusesRepository
}

func NewAddBusUC(db domain.IBusesRepository) *AddBusUC {
	return &AddBusUC{db: db}
}

func (uc *AddBusUC) Run(bus domain.Buses) error {
	return uc.db.Save(bus)
}