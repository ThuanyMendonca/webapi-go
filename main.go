package main

import (
	"github.com/ThuanyMendonca/webapi-go/database"
	"github.com/ThuanyMendonca/webapi-go/server"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {

	tracer.Start(
		tracer.WithService("golang-api"),
		tracer.WithEnv("dev"),
	)
	defer tracer.Stop()

	database.StartDB()
	server := server.NewServer()

	server.Run()
}
