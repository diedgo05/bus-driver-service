package dependencies

import (
	"bus-driver/src/core"
	"bus-driver/src/turnos/application"
	"bus-driver/src/turnos/application/services"
	"bus-driver/src/turnos/infraestructure"
	"bus-driver/src/turnos/infraestructure/http/controllers"
	"fmt"
)

var (
    mySQL        infraestructure.MySQL
    rabbitRepo   *infraestructure.RabbitRepository // adapter  de RabbitMQ
    eventService *services.Event   // service de event
)

func Init() {
    
    db, err := core.InitMySQL()
    if err != nil {
        fmt.Println("Error en el servidor de la base de datos:", err)
        return
    }

    mySQL = *infraestructure.NewMySQL(db)

	rabbitCn, err := core.Connect()
	if err != nil {
		fmt.Println("Error en la conexión a RabbitMQ:", err)
		return
	}

	rabbitRepo = infraestructure.NewRabbitRepository(rabbitCn)

	eventService = services.NewEvent(rabbitRepo)
}

func CrearTurnoCtrl() *controllers.CrearTurnoCtrl {
    ucCrearTurno := application.NewCrearTurnoUC(&mySQL)
    return controllers.NewCrearTurnoCtrl(ucCrearTurno, eventService)
}

// func FinalizarTurnoCtrl() *controllers.FinalizarTurnoCtrl {
// 	ucFinalizarTurno := application.NewFinalizarTurnoUC(&mySQL)
// 	return controllers.NewFinalizarTurnoCtrl(ucFinalizarTurno, eventService)
// }

