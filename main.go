package main

import (
	"fmt"

	"github.com/leomarquezani/rest-api/service"
)

var appName = "accountservice"

func main() {

	fmt.Println("Starting: " + appName)
	service.StartWebServer("6767")
}
