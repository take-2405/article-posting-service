package main

import (
	"net/http"
	_ "net/http/pprof"
	"prac-orm-transaction/config"
	"prac-orm-transaction/di"
	router2 "prac-orm-transaction/presentation/router"
)

func main() {

	//ルーターを初期化
	router := router2.NewServer()
	//ルーティングとDI
	di.InsertUserDI(router)
	port := config.GetServerPort()
	//ルーター起動
	http.ListenAndServe(port, router.Router)
}
