package main

import (
	pb "chika-climate/proto/micro-service-proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type temperatureByYearServer struct {
	pb.UnimplementedClimateDataServiceServer
}

var port = ":5000"

func (s *temperatureByYearServer) GetTemperatureByYear(ctx context.Context, year *pb.Year) (*pb.Temperature, error) {
	return &pb.Temperature{Value: temperatureByYear[year.Value]}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterClimateDataServiceServer(grpcServer, &temperatureByYearServer{})
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

var temperatureByYear = map[int32]int32{
	1990: 2,
	1991: 3,
	1992: 4,
}
