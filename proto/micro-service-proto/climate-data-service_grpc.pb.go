// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package climate_data_service

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

// ClimateDataServiceClient is the client API for ClimateDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClimateDataServiceClient interface {
	GetTemperatureByYear(ctx context.Context, in *Year, opts ...grpc.CallOption) (*Temperature, error)
}

type climateDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClimateDataServiceClient(cc grpc.ClientConnInterface) ClimateDataServiceClient {
	return &climateDataServiceClient{cc}
}

func (c *climateDataServiceClient) GetTemperatureByYear(ctx context.Context, in *Year, opts ...grpc.CallOption) (*Temperature, error) {
	out := new(Temperature)
	err := c.cc.Invoke(ctx, "/chika.climate.ClimateDataService/GetTemperatureByYear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClimateDataServiceServer is the server API for ClimateDataService service.
// All implementations must embed UnimplementedClimateDataServiceServer
// for forward compatibility
type ClimateDataServiceServer interface {
	GetTemperatureByYear(context.Context, *Year) (*Temperature, error)
	mustEmbedUnimplementedClimateDataServiceServer()
}

// UnimplementedClimateDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClimateDataServiceServer struct {
}

func (UnimplementedClimateDataServiceServer) GetTemperatureByYear(context.Context, *Year) (*Temperature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemperatureByYear not implemented")
}
func (UnimplementedClimateDataServiceServer) mustEmbedUnimplementedClimateDataServiceServer() {}

// UnsafeClimateDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClimateDataServiceServer will
// result in compilation errors.
type UnsafeClimateDataServiceServer interface {
	mustEmbedUnimplementedClimateDataServiceServer()
}

func RegisterClimateDataServiceServer(s grpc.ServiceRegistrar, srv ClimateDataServiceServer) {
	s.RegisterService(&ClimateDataService_ServiceDesc, srv)
}

func _ClimateDataService_GetTemperatureByYear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Year)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClimateDataServiceServer).GetTemperatureByYear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chika.climate.ClimateDataService/GetTemperatureByYear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClimateDataServiceServer).GetTemperatureByYear(ctx, req.(*Year))
	}
	return interceptor(ctx, in, info, handler)
}

// ClimateDataService_ServiceDesc is the grpc.ServiceDesc for ClimateDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClimateDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chika.climate.ClimateDataService",
	HandlerType: (*ClimateDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTemperatureByYear",
			Handler:    _ClimateDataService_GetTemperatureByYear_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc-service/climate-data-service.proto",
}