package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handleRequest()
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got the request")
	fmt.Fprintf(w, "hello World")
}

func handleRequest() {
	http.HandleFunc("/", helloWorld)
	log.Fatal(http.ListenAndServe(":9091", nil))
}
