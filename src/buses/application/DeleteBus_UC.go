package application

import "bus-driver/src/buses/domain"

type DeleteBusByIDUC struct {
	db domain.IBusesRepository
}

func NewDeleteBusByIDUC(db domain.IBusesRepository) *DeleteBusByIDUC {
	return &DeleteBusByIDUC{db: db}
}

func (uc *DeleteBusByIDUC) Run(idBus int) error {
	return uc.db.DeleteByID(idBus)
}