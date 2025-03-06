package application

import "bus-driver/src/choferes/domain"

type GetAllChoferesUC struct {
	db domain.IChoferesRepository
}

func NewGetAllChoferesUC(db domain.IChoferesRepository) *GetAllChoferesUC {
	return &GetAllChoferesUC{db: db}
}

func (uc *GetAllChoferesUC) Run() ([]domain.Chofer, error) {
	return uc.db.FindAll()
}