package dependencies

import (
	"bus-driver/src/choferes/application"
	"bus-driver/src/choferes/infraestructure"
	"bus-driver/src/choferes/infraestructure/http/controllers"
	"bus-driver/src/core"
	"fmt"
)

var (
	mySQL infraestructure.MySQL
)

func Init() {
	db, err := core.InitMySQL()
	if err != nil {
		fmt.Println(err)
	}
	mySQL = *infraestructure.NewMySQL(db)
}

func AddChoferCtrl() *controllers.AddChoferCtrl {
	ucAddChofer := application.NewAddChoferUC(&mySQL)
	return controllers.NewAddChoferCtrl(ucAddChofer)
}

func GetChoferesCtrl() *controllers.GetChoferesCtrl {
	ucGetChoferes := application.NewGetAllChoferesUC(&mySQL)
	return controllers.NewGetChoferesCtrl(ucGetChoferes)
}

func UpdateChoferCtrl() *controllers.UpdateByIDChoferController {
	ucUpdateChofer := application.NewUpdateChoferUC(&mySQL)
	return controllers.NewUpdateByIDChoferController(ucUpdateChofer)
}

func DeleteChoferCtrl() *controllers.DeleteChoferCtrl {
	ucDeleteChofer := application.NewDeleteChoferUC(&mySQL)
	return controllers.NewDeleteChoferCtrl(ucDeleteChofer)
}

func GetChoferByIDCtrl() *controllers.GetChoferByIDCtrl {
	ucGetChoferByID := application.NewGetChoferByIDUC(&mySQL)
	return controllers.NewGetChoferByIDCtrl(ucGetChoferByID)
}
