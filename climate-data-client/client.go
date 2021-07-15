package main

import (
	"context"
	"log"
	"time"

	pb "chika-climate/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:5000"
)

func getTemperatureByYear(client pb.ClimateDataServiceClient, year *pb.Year) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	temperature, err := client.GetTemperatureByYear(ctx, year)
	//log.Println(year)
	if err != nil {
		log.Fatalf("GetTemperature fails", err)
	}
	log.Println(temperature)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect")
	}
	defer conn.Close()
	c := pb.NewClimateDataServiceClient(conn)
	getTemperatureByYear(c, &pb.Year{Value: 1990})
}
