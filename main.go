package main

import (
	"log"
	"net/http"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("listening on port 3000...")
	http.ListenAndServe(":3000", nil)
}
