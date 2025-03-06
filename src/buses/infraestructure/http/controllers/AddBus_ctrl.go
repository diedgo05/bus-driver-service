package controllers

import (
	"bus-driver/src/buses/application"
	"bus-driver/src/buses/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddBusCtrl struct {
	uc *application.AddBusUC
}

func NewAddBusCtrl(uc *application.AddBusUC) *AddBusCtrl {
	return &AddBusCtrl{uc: uc}
}

func (ctrl *AddBusCtrl) Run(c *gin.Context) {
	var buses domain.Buses

	if err := c.ShouldBindJSON(&buses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	}

	err := ctrl.uc.Run(buses)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": true,
			"data": gin.H{
				"type":  "buses",
				"idBus": buses.IdBus,
				"attributes": gin.H{
					"placa":     buses.Placa,
					"capacidad": buses.Capacidad,
					"chofer":    buses.ChoferID,
				},
			},
		})
	}
}