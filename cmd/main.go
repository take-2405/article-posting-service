package main

import (
	"log"
	"net/http"
	"prac-orm-transaction/api/middleware"
	router2 "prac-orm-transaction/api/router"
)

func main() {
	if err := middleware.Init(); err != nil {
		log.Fatal(err)
	}

	router := router2.NewServer()
	router.Routing()

	//if *router. {
	//	// fmt.Println(docgen.JSONRoutesDoc(r))
	//	fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
	//		ProjectPath: "github.com/go-chi/chi/v5",
	//		Intro:       "Welcome to the chi/_examples/rest generated docs.",
	//	}))
	//	return
	//}

	http.ListenAndServe(":8000", router.Router)
}
