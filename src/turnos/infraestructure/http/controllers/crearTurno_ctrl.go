package controllers

import (
	"bus-driver/src/turnos/application"
	"bus-driver/src/turnos/application/services"
	"bus-driver/src/turnos/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CrearTurnoCtrl struct {
	uc    *application.CrearTurnoUC
	event *services.Event
}

func NewCrearTurnoCtrl(uc *application.CrearTurnoUC, event *services.Event) *CrearTurnoCtrl {
	return &CrearTurnoCtrl{uc: uc, event: event}
}

func (ctrl *CrearTurnoCtrl) Run(c *gin.Context) {
	var turno domain.Turno

	// Bind JSON a la estructura Turno
	if err := c.ShouldBindJSON(&turno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Guardar el turno usando el caso de uso
	if err := ctrl.uc.Run(turno); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar el turno: " + err.Error()})
		return
	}

	// Publicar el evento en RabbitMQ
	if err := ctrl.event.PublishEvent(turno); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo publicar el evento: " + err.Error()})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusCreated, gin.H{
		"message": "Turno creado exitosamente",
		"data":    turno,
	})
}