package main

import (
	"fmt"
	"log"

	"github.com/nitoc-ict/fushigidane-server/insertlocation"
	"github.com/nitoc-ict/fushigidane-server/mapsapi"
	"github.com/nitoc-ict/fushigidane-server/rdbms"
	"github.com/nitoc-ict/fushigidane-server/router"
)

func main() {
	rdbms.InitDBClient()
	mapsapi.InitMapsAPIClient()
	router := router.NewRouter()

	a, err := insertlocation.ConvertCoordinateToAddress(26.407455, 127.734630)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(a)

	router.Logger.Fatal(router.Start(":5000"))
}
