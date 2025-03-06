package application

import "bus-driver/src/choferes/domain"

type DeleteChoferUC struct {
	db domain.IChoferesRepository
}

func NewDeleteChoferUC(db domain.IChoferesRepository) *DeleteChoferUC {
	return &DeleteChoferUC{db: db}
}

func (uc *DeleteChoferUC) Run(id int) error {
	return uc.db.DeleteByID(id)
}