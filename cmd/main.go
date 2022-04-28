package main

import (
	"log"
	"net"

	"github.com/vaibhavsatam/sample/pkg/v1/api"
	"github.com/vaibhavsatam/sample/pkg/v2/service"
	"google.golang.org/grpc"
)

func main() {
	ls, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	server := grpc.NewServer()

	todoService := service.NewToDoServiceServer()
	api.RegisterTodoServiceServer(server, todoService)

	sensorService := service.NewSensorServiceServer()
	api.RegisterSensorServiceServer(server, sensorService)

	if err := server.Serve(ls); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
