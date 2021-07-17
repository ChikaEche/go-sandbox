package main

import (
	pb "chika-climate/proto/micro-service-proto"
	"context"
	"fmt"
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
	fmt.Println("serrr")
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

var temperatureByYear = map[int32]float64{
	1990: 0.45,
	1991: 0.41,
	1992: 0.22,
	1993: 0.23,
	1994: 0.32,
	1995: 0.45,
	1996: 0.33,
	1997: 0.47,
	1998: 0.61,
	1999: 0.39,
	2000: 0.40,
	2001: 0.54,
	2002: 0.63,
	2003: 0.62,
	2004: 0.54,
	2005: 0.68,
	2006: 0.64,
	2007: 0.66,
	2008: 0.54,
	2009: 0.66,
	2010: 0.72,
	2011: 0.61,
	2012: 0.65,
	2013: 0.68,
	2014: 0.74,
	2015: 0.90,
	2016: 1.01,
	2017: 0.92,
	2018: 0.85,
	2019: 0.98,
	2020: 1.02,
}
