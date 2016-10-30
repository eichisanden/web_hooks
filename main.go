package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}
