package main

import (
	"fmt"
	"net/http"
)

func getTemperatureByYear(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome to my API!\n")
}

func main() {
	http.HandleFunc("/", getTemperatureByYear)
	http.ListenAndServe(":3000", nil)
}
