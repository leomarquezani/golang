package main

import (
	"fmt"

	"github.com/leomarquezani/rest-api/dbclient"
	"github.com/leomarquezani/rest-api/service"
)

var appName = "accountservice"

func main() {

	fmt.Println("Starting: " + appName)
	initializeBoltClient()
	service.StartWebServer("6767")
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
