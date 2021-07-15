package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	apiPb "chika-climate/proto/api-client-proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8000"
)

func getTemperatureByYear(w http.ResponseWriter, req *http.Request) {
	year, intErr := strconv.ParseInt(req.URL.Query()["year"][0], 10, 32)
	conn, dialErr := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if dialErr != nil {
		log.Fatalf("did not connect")
	}
	defer conn.Close()
	client := apiPb.NewApiProtoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if intErr != nil {
		panic(intErr)
	}
	temperature, err := client.GetTemperature(ctx, &apiPb.Yr{Value: int32(year)})
	if err != nil {
		log.Fatalf("GetTemperature fails", err)
	}
	fmt.Fprintln(w, temperature)
}

func main() {
	http.HandleFunc("/", getTemperatureByYear)
	http.ListenAndServe(":3000", nil)
}
