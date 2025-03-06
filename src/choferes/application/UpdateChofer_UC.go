package application

import "bus-driver/src/choferes/domain"

type UpdateChoferUC struct {
	db domain.IChoferesRepository
}

func NewUpdateChoferUC(db domain.IChoferesRepository) *UpdateChoferUC {
	return &UpdateChoferUC{db: db}
}

func (uc *UpdateChoferUC) Run(id int, chofer domain.Chofer) error {
	return uc.db.UpdateByID(id, chofer)
}