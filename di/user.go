package di

import (
	"prac-orm-transaction/infrastructure"
	"prac-orm-transaction/interface/controller"
	"prac-orm-transaction/interface/router"
	"prac-orm-transaction/usecase"
)

func InsertUserDI(router *router.Server) {
	conn := infrastructure.NewMysqlRepository()

	userQuery := infrastructure.NewUserPersistence(*conn)
	articleQuery := infrastructure.NewArticlePersistence(*conn)

	userUseCase := usecase.NewUserUseCase(userQuery)
	articleUseCase := usecase.NewArticleUseCase(articleQuery)

	useHandler := controller.NewUserHandler(userUseCase)
	articleHandler := controller.NewArticleHandler(articleUseCase)
	router.Routing(useHandler, articleHandler)
}
