package main

import (
	"net/http"
	router2 "prac-orm-transaction/api/router"
	"prac-orm-transaction/di"
)

func main() {
	//ルーターを初期化
	router := router2.NewServer()
	//ルーティングとDI
	di.InsertUserDI(router)
	//if *router. {
	//	// fmt.Println(docgen.JSONRoutesDoc(r))
	//	fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
	//		ProjectPath: "github.com/go-chi/chi/v5",
	//		Intro:       "Welcome to the chi/_examples/rest generated docs.",
	//	}))
	//	return
	//}

	//ルーター起動
	http.ListenAndServe(":8000", router.Router)
}
