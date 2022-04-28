package service

import (
	"context"

	"github.com/vaibhavsatam/sample/pkg/v1/api"
)

type sensorServiceServer struct {
	api.UnimplementedSensorServiceServer
}

func NewSensorServiceServer() api.SensorServiceServer {
	return &sensorServiceServer{}
}

// TempSensor implements api.SensorServiceServer
func (*sensorServiceServer) TempSensor(context.Context, *api.TempRequest) (*api.TempResponse, error) {
	return &api.TempResponse{
		Value: 2,
	}, nil
}
