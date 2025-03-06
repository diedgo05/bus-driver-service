package application

import "bus-driver/src/choferes/domain"

type GetChoferByIDUC struct {
	db domain.IChoferesRepository
}

func NewGetChoferByIDUC(db domain.IChoferesRepository) *GetChoferByIDUC {
	return &GetChoferByIDUC{db: db}
}

func (uc *GetChoferByIDUC) Run(id int) ([]domain.Chofer, error) {
	return uc.db.FindByID(id)
}