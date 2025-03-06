package dependencies

import (
	"bus-driver/src/buses/application"
	"bus-driver/src/buses/infraestructure"
	"bus-driver/src/buses/infraestructure/http/controllers"
	"bus-driver/src/core"
)

var (
	mySQL infraestructure.MySQL
)

func Init() {
	db, err := core.InitMySQL()
	if err != nil {
		return
	}
	mySQL = *infraestructure.NewMySQL(db)
}

func AddBusCtrl() *controllers.AddBusCtrl {
	ucAddBus := application.NewAddBusUC(&mySQL)
	return controllers.NewAddBusCtrl(ucAddBus)
}

func GetBusesCtrl() *controllers.GetBusesCtrl {
	ucGetBuses := application.NewGetBusesUC(&mySQL)
	return controllers.NewGetBusesCtrl(ucGetBuses)
}

func DeleteBusCtrl() *controllers.DeleteBusCtrl {
	ucDeleteBus := application.NewDeleteBusByIDUC(&mySQL)
	return controllers.NewDeleteBusCtrl(ucDeleteBus)
}

func UpdateBusCtrl() *controllers.UpdateBusCtrl {
	ucUpdateBus := application.NewUpdateBusByIDUC(&mySQL)
	return controllers.NewUpdateBusCtrl(ucUpdateBus)
}

func GetBusByIDCtrl() *controllers.GetBusByIDCtrl {
	ucGetBusByID := application.NewGetBusByIDUC(&mySQL)
	return controllers.NewGetBusByIDCtrl(ucGetBusByID)
}

