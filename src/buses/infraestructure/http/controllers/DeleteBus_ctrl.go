package controllers

import (
	"bus-driver/src/buses/application"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type DeleteBusCtrl struct {
	uc *application.DeleteBusByIDUC
}

func NewDeleteBusCtrl(uc *application.DeleteBusByIDUC) *DeleteBusCtrl {
	return &DeleteBusCtrl{uc: uc}
}

func (ctrl *DeleteBusCtrl) Run(c *gin.Context) {
	id := c.Param("idBus")
	busID, err := strconv.Atoi(id)	

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser un número"})
		return
	}

	err = ctrl.uc.Run(busID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Bus eliminado correctamente"})
	}
}