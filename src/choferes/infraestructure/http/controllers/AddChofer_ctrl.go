package controllers

import (
	"bus-driver/src/choferes/application"
	"bus-driver/src/choferes/domain"
	"net/http"
	"github.com/gin-gonic/gin"
)

type AddChoferCtrl struct {
	uc *application.AddChoferUC
}

//Constructor de AddChoferController
func NewAddChoferCtrl(uc *application.AddChoferUC) *AddChoferCtrl {
	return &AddChoferCtrl{uc: uc}
}

func (ctrl *AddChoferCtrl) Run(c *gin.Context) {
	var choferes domain.Chofer

	if err := c.ShouldBindJSON(&choferes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	} 

	err := ctrl.uc.Run(choferes)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": true,
			"data": gin.H{
				"type": "choferes",
				"idChofer": choferes.ID,
				"attributes": gin.H{
					"nombre": choferes.Nombre,
					"apellidos":choferes.Apellido_p + choferes.Apellido_m,
					"edad":choferes.Edad,
				},
			},
		})
	}
}