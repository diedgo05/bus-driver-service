package application

import (
	"bus-driver/src/turnos/domain"
)

type CrearTurnoUC struct {
	repo domain.ITurnoRepository
}

func NewCrearTurnoUC(repo domain.ITurnoRepository) *CrearTurnoUC {
	return &CrearTurnoUC{repo: repo}
}

func (uc *CrearTurnoUC) Run(turno domain.Turno) error {
	return uc.repo.CrearTurno(turno)
}