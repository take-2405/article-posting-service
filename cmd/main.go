package main

import (
	"net/http"
	"prac-orm-transaction/di"
	router2 "prac-orm-transaction/interface/router"
)

func main() {
	//ルーターを初期化
	router := router2.NewServer()
	//ルーティングとDI
	di.InsertUserDI(router)

	//ルーター起動
	http.ListenAndServe(":8000", router.Router)
}
