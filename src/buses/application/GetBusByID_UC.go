package application

import "bus-driver/src/buses/domain"

type GetBusByIDUC struct {
	db domain.IBusesRepository
}

func NewGetBusByIDUC(db domain.IBusesRepository) *GetBusByIDUC {
	return &GetBusByIDUC{db: db}
}

func (uc *GetBusByIDUC) Run(choferID int) ([]domain.Buses, error) {
	return uc.db.FindBusByIdChofer(choferID)
}