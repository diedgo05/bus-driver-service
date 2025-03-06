package routes

import (
	"bus-driver/src/choferes/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	routes := router.Group("/choferes")

	addChofer := dependencies.AddChoferCtrl()
	getAllChofer := dependencies.GetChoferesCtrl()
	updateChofer := dependencies.UpdateChoferCtrl()
	deleteChofer := dependencies.DeleteChoferCtrl()
	getChoferByID := dependencies.GetChoferByIDCtrl()

	routes.POST("/", addChofer.Run)
	routes.GET("/", getAllChofer.Run)
	routes.GET("/:id", getChoferByID.Run)
	routes.PUT("/:id", updateChofer.Run)
	routes.DELETE("/:id", deleteChofer.Run)
}