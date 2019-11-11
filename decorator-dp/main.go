package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Starting http server")
	http.HandleFunc("/", handleRequest)

	//logHeader decorator function adds functionality to handleRequest
	http.HandleFunc("/header", logHeader(handleRequest))

	http.ListenAndServe(":3000", nil)

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<xml>test</xml>")
}

func logHeader(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("decorating request")
		fmt.Printf("Received request %v", r.Header)
		fn(w, r)
	}

}
