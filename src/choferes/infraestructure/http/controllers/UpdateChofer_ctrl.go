package controllers

import (
	"bus-driver/src/choferes/application"
	"bus-driver/src/choferes/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateByIDChoferController struct {
	uc *application.UpdateChoferUC
}

func NewUpdateByIDChoferController(uc *application.UpdateChoferUC) *UpdateByIDChoferController {
	return &UpdateByIDChoferController{uc: uc}
}

func (ctrl *UpdateByIDChoferController) Run(c *gin.Context) {
	var chofer domain.Chofer
	if err := c.ShouldBindJSON(&chofer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	}

	idChofer := c.Param("id")
	id, err := strconv.Atoi(idChofer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser un número"})
		return
	}
	if err := ctrl.uc.Run(id, chofer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} 
		c.JSON(http.StatusOK, gin.H{"message": "Chofer actualizado correctamente"})
}