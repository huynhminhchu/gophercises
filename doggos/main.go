package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to chudeptrai site!</h1>")
}
func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server on :80....")
	http.ListenAndServe(":3000", nil)
}
