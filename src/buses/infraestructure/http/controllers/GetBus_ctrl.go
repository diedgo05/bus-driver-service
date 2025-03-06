package controllers

import (
	"bus-driver/src/buses/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetBusesCtrl struct {
	uc *application.GetBusesUC
}

func NewGetBusesCtrl(uc *application.GetBusesUC) *GetBusesCtrl {
	return &GetBusesCtrl{uc: uc}
}

func (ctrl *GetBusesCtrl) Run(c *gin.Context) {
	buses, err := ctrl.uc.Run()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":"Todos los campos son requeridos",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data":   buses,
		})
	}
}

