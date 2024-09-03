package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//serve static files
	fs := http.FileServer(http.Dir("./static"))             // file server that serves file from the dir
	http.Handle("/static", http.StripPrefix("/static", fs)) // sets route, route static will be handled. stripeprefix removes static prefix before file server finds it

	// handle root route
	http.HandleFunc("/", homeHandler)

	// start server
	fmt.Println("Starting server on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start the server")
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/index.html")
}
