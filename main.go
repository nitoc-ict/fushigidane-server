package main

import "github.com/nitoc-ict/fushigidane-server/router"

func main() {
	router := router.NewRouter()

	router.Logger.Fatal(router.Start(":5000"))
}
