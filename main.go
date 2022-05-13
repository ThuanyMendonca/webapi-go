package main

import (
	"github.com/ThuanyMendonca/webapi-go/database"
	"github.com/ThuanyMendonca/webapi-go/server"
)

func main() {

	database.StartDB()
	server := server.NewServer()

	server.Run()
}
