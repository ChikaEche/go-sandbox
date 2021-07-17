package main

import (
	pb "chika-climate/proto/micro-service-proto"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8000"
)

func getTemperatureByYear(w http.ResponseWriter, req *http.Request) {
	year, intErr := strconv.ParseInt(req.URL.Query()["year"][0], 10, 32)
	conn, dialErr := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	fmt.Println("here", conn.GetState(), dialErr)
	if dialErr != nil {
		log.Fatalf("did not connect")
	}
	defer conn.Close()
	client := pb.NewClimateDataServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if intErr != nil {
		panic(intErr)
	}
	temperature, err := client.GetTemperatureByYear(ctx, &pb.Year{Value: int32(year)})
	if err != nil {
		log.Fatalf("GetTemperature fails", err)
	}
	fmt.Fprintln(w, temperature)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	clientCmd := exec.Command("clent.exe")
	clientErr := clientCmd.Start()
	if clientErr != nil {
		fmt.Println("Failed to start client server")
	}

	serverCmd := exec.Command("server.exe")
	serverErr := serverCmd.Start()
	if serverErr != nil {
		fmt.Println("Failed to start server")
	}
	http.HandleFunc("/", getTemperatureByYear)
	http.ListenAndServe(":"+port, nil)
}
