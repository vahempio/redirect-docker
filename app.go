package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var port string = "80"

func main() {
	if os.Getenv("URL") == "" {
		fmt.Println("$URL not defined, use default")
		os.Setenv("URL", "http://foo.bar.localhost")
	}
	fmt.Println("Starting up on port "+port)
	fmt.Println("Redirect requests to "+os.Getenv("URL"))
	http.HandleFunc("/", redirect)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func redirect(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Location", os.Getenv("URL"))
	w.WriteHeader(301)
}
