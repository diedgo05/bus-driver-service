package routes

import (
	"bus-driver/src/buses/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group(("/buses"))

	addBus := dependencies.AddBusCtrl()
	getAllBus := dependencies.GetBusesCtrl()
	updateBus := dependencies.UpdateBusCtrl()
	getBusByIdChofer := dependencies.GetBusByIDCtrl()
	deleteBus := dependencies.DeleteBusCtrl()

	routes.POST("/", addBus.Run)
	routes.GET("/", getAllBus.Run)
	routes.PUT("/:idBus", updateBus.Run)
	routes.GET("/:choferID", getBusByIdChofer.Run)
	routes.DELETE("/:idBus", deleteBus.Run)
}