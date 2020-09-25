package main

import (
	"github.com/nitoc-ict/fushigidane-server/mapsapi"
	"github.com/nitoc-ict/fushigidane-server/rdbms"
	"github.com/nitoc-ict/fushigidane-server/router"
)

func main() {
	rdbms.InitDBClient()
	mapsapi.InitMapsAPIClient()
	router := router.NewRouter()

	router.Logger.Fatal(router.Start(":5000"))
}
