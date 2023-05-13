package main

import (
	"encoding/json"
	"fmt"
	"http-apis/models"
	"log"
	"net/http"
)

func main() {
	handleRequest()
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	a := models.Article{
		Title: "Hello",
	}
	fmt.Println("Got the request: %w", a)
	fmt.Fprintf(w, "hello World")
}

func articles(w http.ResponseWriter, r *http.Request) {
	allArticles := models.GetArticles()
	fmt.Println("Articles endpoint is hit")
	json.NewEncoder(w).Encode(allArticles)
}

func handleRequest() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/articles", articles)
	log.Fatal(http.ListenAndServe(":9091", nil))
}
