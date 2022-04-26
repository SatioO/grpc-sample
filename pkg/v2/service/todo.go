package service

import (
	"context"

	todo "github.com/vaibhavsatam/sample/pkg/v1/api"
)

type toDoServiceServer struct {
	todo.UnimplementedTodoServiceServer
}

func NewToDoServiceServer() todo.TodoServiceServer {
	return &toDoServiceServer{}
}

// Create implements todo.TodoServiceServer
func (*toDoServiceServer) Create(context.Context, *todo.CreateTodoRequest) (*todo.CreateTodoResponse, error) {
	panic("unimplemented")
}
