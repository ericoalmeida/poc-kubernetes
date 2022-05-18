package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/load-config", LoadConfig)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	USER_NAME := os.Getenv("USER_NAME")

	fmt.Fprintf(w, "<h1> Hello %s!!! <h1>", USER_NAME)
}

func LoadConfig(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/go/test/.env")

	if err != nil {
		log.Fatal("Error loading file", err)
	}

	fmt.Fprintf(w, "File data: %s", string(data))
}
