package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	log.Println("HandleFunc hello")
	fmt.Fprintf(w, "hello\n")
}

func main() {
	log.Println("running on port 8080")

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)

}
