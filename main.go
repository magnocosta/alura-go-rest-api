package main

import (
	"go-rest-api/database"
	"go-rest-api/routes"
)

func main() {
	database.OpenDBConnection()
	routes.Init()
}
