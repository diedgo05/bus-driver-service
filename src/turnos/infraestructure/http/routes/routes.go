package routes

import (
	"bus-driver/src/turnos/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/turnos")

	crearTurno := dependencies.CrearTurnoCtrl()
	// finalizarTurno := dependencies.FinalizarTurnoCtrl()

	routes.POST("/", crearTurno.Run)
	// routes.POST("/:id", finalizarTurno.Run)
}