package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := "8090"

	http.HandleFunc("/test", rTest)
	http.HandleFunc("/", rNone)

	fmt.Printf("Starting http server on port %v... \n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func rNone(w http.ResponseWriter, req *http.Request) {
	log("/")
	fmt.Fprintf(w, "hello from /\n")
}

func rTest(w http.ResponseWriter, req *http.Request) {
	log("/test")
	fmt.Fprintf(w, "hello from /test\n")
}

func log(route string) {
	fmt.Printf("Request recieved on %s\n", route)
}
