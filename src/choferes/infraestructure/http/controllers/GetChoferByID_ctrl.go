package controllers

import (
	"bus-driver/src/choferes/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetChoferByIDCtrl struct {
	uc *application.GetChoferByIDUC
}

func NewGetChoferByIDCtrl(uc *application.GetChoferByIDUC) *GetChoferByIDCtrl{
	return &GetChoferByIDCtrl{uc:uc}
}

func (ctrl *GetChoferByIDCtrl) Run(c *gin.Context) {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser un número"})
		return
	}

	chofer, err := ctrl.uc.Run(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"chofer encontrado por id": chofer})
	}
}