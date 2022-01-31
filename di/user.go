package di

import (
	"prac-orm-transaction/infrastructure"
	"prac-orm-transaction/presentation/controller"
	"prac-orm-transaction/presentation/router"
	"prac-orm-transaction/usecase"
)

func InsertUserDI(router *router.Server) {
	conn := infrastructure.NewMysqlRepository()

	userQuery := infrastructure.NewUserPersistence(*conn)
	articleQuery := infrastructure.NewArticlePersistence(*conn)

	userUseCase := usecase.NewAuthUseCase(userQuery)
	articleUseCase := usecase.NewArticleUseCase(articleQuery)

	useHandler := controller.NewUserHandler(userUseCase)
	articleHandler := controller.NewArticleHandler(articleUseCase)
	router.Routing(useHandler, articleHandler)
}
