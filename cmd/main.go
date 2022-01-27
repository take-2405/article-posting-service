package main

import (
	"log"
	"prac-orm-transaction/api/middleware"
)

func main() {
	if err := middleware.Init(); err != nil {
		log.Fatal(err)
	}
}
