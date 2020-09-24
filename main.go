package main

import (
	"github.com/nitoc-ict/fushigidane-server/rdbms"
	"github.com/nitoc-ict/fushigidane-server/router"
)

func main() {
	rdbms.InitDBClient()
	router := router.NewRouter()

	router.Logger.Fatal(router.Start(":5000"))
}
