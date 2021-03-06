// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/v1/proto/sensor.service.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SensorServiceClient is the client API for SensorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SensorServiceClient interface {
	TempSensor(ctx context.Context, in *TempRequest, opts ...grpc.CallOption) (*TempResponse, error)
}

type sensorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSensorServiceClient(cc grpc.ClientConnInterface) SensorServiceClient {
	return &sensorServiceClient{cc}
}

func (c *sensorServiceClient) TempSensor(ctx context.Context, in *TempRequest, opts ...grpc.CallOption) (*TempResponse, error) {
	out := new(TempResponse)
	err := c.cc.Invoke(ctx, "/api.SensorService/TempSensor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SensorServiceServer is the server API for SensorService service.
// All implementations must embed UnimplementedSensorServiceServer
// for forward compatibility
type SensorServiceServer interface {
	TempSensor(context.Context, *TempRequest) (*TempResponse, error)
	mustEmbedUnimplementedSensorServiceServer()
}

// UnimplementedSensorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSensorServiceServer struct {
}

func (UnimplementedSensorServiceServer) TempSensor(context.Context, *TempRequest) (*TempResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TempSensor not implemented")
}
func (UnimplementedSensorServiceServer) mustEmbedUnimplementedSensorServiceServer() {}

// UnsafeSensorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SensorServiceServer will
// result in compilation errors.
type UnsafeSensorServiceServer interface {
	mustEmbedUnimplementedSensorServiceServer()
}

func RegisterSensorServiceServer(s grpc.ServiceRegistrar, srv SensorServiceServer) {
	s.RegisterService(&SensorService_ServiceDesc, srv)
}

func _SensorService_TempSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TempRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorServiceServer).TempSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SensorService/TempSensor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorServiceServer).TempSensor(ctx, req.(*TempRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SensorService_ServiceDesc is the grpc.ServiceDesc for SensorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SensorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.SensorService",
	HandlerType: (*SensorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TempSensor",
			Handler:    _SensorService_TempSensor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/proto/sensor.service.proto",
}
