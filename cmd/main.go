package main

import (
	"log"
	"net"

	todo "github.com/vaibhavsatam/sample/pkg/v1/api"
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
	todo.RegisterTodoServiceServer(server, todoService)

	if err := server.Serve(ls); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
