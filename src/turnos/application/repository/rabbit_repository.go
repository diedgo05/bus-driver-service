package repository

import "bus-driver/src/turnos/domain"

type IRabbitRepository interface {
	PublishMessage(turno domain.Turno) error
}