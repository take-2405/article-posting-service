package di

import (
	"prac-orm-transaction/api/controller"
	"prac-orm-transaction/api/router"
	"prac-orm-transaction/application"
	"prac-orm-transaction/infrastructure"
)

func InsertUserDI(router *router.Server) {
	conn := infrastructure.NewMysqlRepository()
	userQuery := infrastructure.NewUserPersistence(*conn)
	userUseCase := application.NewUserUseCase(userQuery)
	useHandler := controller.NewUserHandler(userUseCase)
	router.Routing(useHandler)
}
