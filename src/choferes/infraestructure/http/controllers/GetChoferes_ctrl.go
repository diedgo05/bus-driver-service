package controllers

import (
	"bus-driver/src/choferes/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetChoferesCtrl struct {
	uc *application.GetAllChoferesUC
}

func NewGetChoferesCtrl(uc *application.GetAllChoferesUC) *GetChoferesCtrl {
	return &GetChoferesCtrl{uc: uc}
}

func (ctrl *GetChoferesCtrl) Run(c *gin.Context) {
	choferes, err := ctrl.uc.Run()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"choferes": choferes})
}