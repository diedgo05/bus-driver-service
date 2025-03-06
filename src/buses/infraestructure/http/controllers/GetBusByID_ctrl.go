package controllers

import (
	"bus-driver/src/buses/application"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
)

type GetBusByIDCtrl struct {
	uc *application.GetBusByIDUC
}

func NewGetBusByIDCtrl(uc *application.GetBusByIDUC) *GetBusByIDCtrl {
	return &GetBusByIDCtrl{uc: uc}
}

func (ctrl *GetBusByIDCtrl) Run(c *gin.Context) {
	id := c.Param("choferID")
	ChoferID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser un número"})
		return 
	}

	buses, err := ctrl.uc.Run(ChoferID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"buses": buses})
	}
}