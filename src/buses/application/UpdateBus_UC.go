package application

import "bus-driver/src/buses/domain"

type UpdateBusByIDUC struct {
	db domain.IBusesRepository
}

func NewUpdateBusByIDUC(db domain.IBusesRepository) *UpdateBusByIDUC {
	return &UpdateBusByIDUC{db: db}
}

func (uc *UpdateBusByIDUC) Run(idBus int, bus domain.Buses) error {
	return uc.db.UpdateByID(idBus, bus)
}
