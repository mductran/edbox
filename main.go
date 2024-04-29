package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/hello", hello)

	log.Fatal(http.ListenAndServe(":3333", mux))
}
