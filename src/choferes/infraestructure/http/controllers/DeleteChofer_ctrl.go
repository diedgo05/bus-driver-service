package controllers

import (
	"bus-driver/src/choferes/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteChoferCtrl struct {
	uc *application.DeleteChoferUC
}

func NewDeleteChoferCtrl(uc *application.DeleteChoferUC) *DeleteChoferCtrl {
	return &DeleteChoferCtrl{uc: uc}
}

func (ctrl *DeleteChoferCtrl) Run(c *gin.Context) {
	id := c.Param("id")
	choferID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser un número"})
		return
	}

	err = ctrl.uc.Run(choferID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Chofer eliminado correctamente"})
	}
}