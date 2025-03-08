package services

import (
	"bus-driver/src/turnos/application/repository"
	"bus-driver/src/turnos/domain"
)

type Event struct {
	rabbitRepository repository.IRabbitRepository
}

func NewEvent (rabbitRepository repository.IRabbitRepository) *Event {
	return &Event{rabbitRepository: rabbitRepository}
}

func (s *Event) PublishEvent(turno domain.Turno) error {
	return s.rabbitRepository.PublishMessage(turno)
}