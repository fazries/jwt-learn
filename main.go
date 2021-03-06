package main

import (
	"log"
	"net/http"
)

func main() {
	// "Login" and "Welcome" are the handler that will be implement
	http.HandleFunc("/login", Login)
	http.HandleFunc("/welcome", Welcome)
	//http.HandleFunc("/refresh", Refresh)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
