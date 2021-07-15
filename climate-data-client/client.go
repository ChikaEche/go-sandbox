package main

import (
	"context"
	"log"
	"net"
	"time"

	apiPb "chika-climate/proto/api-client-proto"
	pb "chika-climate/proto/micro-service-proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:5000"
	port    = ":8000"
)

type temperatureServer struct {
	apiPb.UnimplementedApiProtoServer
}

func getTemperatureByYear(client pb.ClimateDataServiceClient, year *pb.Year) (*apiPb.Temp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	temperature, err := client.GetTemperatureByYear(ctx, year)
	//log.Println(year)
	if err != nil {
		log.Fatalf("GetTemperature fails", err)
	}
	return &apiPb.Temp{Value: temperature.Value}, nil
}

func main() {

	listen, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	apiPb.RegisterApiProtoServer(server, &temperatureServer{})

	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *temperatureServer) GetTemperature(ctx context.Context, year *apiPb.Yr) (*apiPb.Temp, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect")
	}
	defer conn.Close()
	c := pb.NewClimateDataServiceClient(conn)
	return getTemperatureByYear(c, &pb.Year{Value: year.Value})
}
