package main 

import (
    dependenciesBus "bus-driver/src/buses/infraestructure/dependencies"
    routesBus "bus-driver/src/buses/infraestructure/http/routes"
    dependenciesChofer "bus-driver/src/choferes/infraestructure/dependencies"
    routesChofer "bus-driver/src/choferes/infraestructure/http/routes"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    r := gin.Default()
    r.Use(cors.Default())

    // Inicializar choferes
    dependenciesChofer.Init()
    routesChofer.SetupRoutes(r)

    // Inicializar buses
    dependenciesBus.Init()
    routesBus.Routes(r)

    r.Run()
}