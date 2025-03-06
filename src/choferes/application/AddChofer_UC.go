package application

import "bus-driver/src/choferes/domain"

type AddChoferUC struct {
	db domain.IChoferesRepository
}

func NewAddChoferUC (db domain.IChoferesRepository) *AddChoferUC  {
	return &AddChoferUC{db: db}
}

func (uc *AddChoferUC ) Run(chofer domain.Chofer) error {
	return uc.db.Save(chofer)
}
