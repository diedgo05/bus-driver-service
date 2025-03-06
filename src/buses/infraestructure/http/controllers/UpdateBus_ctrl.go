package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"bus-driver/src/buses/application"
	"bus-driver/src/buses/domain"
)

type UpdateBusCtrl struct {
	uc *application.UpdateBusByIDUC
}

func NewUpdateBusCtrl(uc *application.UpdateBusByIDUC) *UpdateBusCtrl {
	return &UpdateBusCtrl{uc: uc}
}

func (ctrl *UpdateBusCtrl) Run(c *gin.Context) {
	var bus domain.Buses
	if err := c.ShouldBindJSON(&bus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Todos los campos son requeridos"})
		return
	}

	idBus := c.Param("idBus")
	id, err := strconv.Atoi(idBus)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser un número"})
		return
	}

	if err := ctrl.uc.Run(id, bus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
 	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Bus actualizado correctamente"})
	}
}