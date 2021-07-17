package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "chika-climate/proto/micro-service-proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:5000"
	port    = ":8000"
)

type temperatureByYearServer struct {
	pb.UnimplementedClimateDataServiceServer
}

func getTemperatureByYear(client pb.ClimateDataServiceClient, year *pb.Year) (*pb.Temperature, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	temperature, err := client.GetTemperatureByYear(ctx, year)
	//log.Println(year)
	if err != nil {
		log.Fatalf("GetTemperature fails", err)
	}
	return &pb.Temperature{Value: temperature.Value}, nil
}

func main() {
	fmt.Println("yoooo")

	listen, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterClimateDataServiceServer(server, &temperatureByYearServer{})

	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *temperatureByYearServer) GetTemperatureByYear(ctx context.Context, year *pb.Year) (*pb.Temperature, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect")
	}
	defer conn.Close()
	c := pb.NewClimateDataServiceClient(conn)
	return getTemperatureByYear(c, year)
}
