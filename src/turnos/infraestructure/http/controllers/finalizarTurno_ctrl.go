package controllers

// import (
// 	"bus-driver/src/turnos/application"
// 	"strconv"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// type FinalizarTurnoCtrl struct {
// 	uc *application.FinalizarTurnoUC
// }

// func NewFinalizarTurnoCtrl(uc *application.FinalizarTurnoUC) *FinalizarTurnoCtrl {
// 	return &FinalizarTurnoCtrl{uc: uc}
// }

// func (ctrl *FinalizarTurnoCtrl) Run(c *gin.Context)  {
// 	id := c.Param("id")
// 	ID, err := strconv.Atoi(id)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return 
// 	}
// 	turno := ctrl.uc.Run(ID, time.Now())
// 	c.JSON(200, gin.H{"turno finalizado": turno})
// }