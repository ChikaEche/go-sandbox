package main

import (
	"fmt"
	"net/http"
)

var c = map[int]float32{
	1960: -0.03,
	1962: 0.03,
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println(c)
	fmt.Fprint(w, c, r.URL.Path[1:])
}
