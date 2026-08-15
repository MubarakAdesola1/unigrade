package main

import (
	"fmt"
	"net/http"
	"os"
)

var subjects = []string{
	"Mathematics",
	"English",
	"Physics",
	"Chemistry",
	"Biology",
	"Computer Science",
}

var units = []int{3, 3, 2, 2, 2, 1}

func main() {
	// Serve static files from the "static" folder
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", showForm)
	http.HandleFunc("/calculate", calculate)

	// Get port from environment variable (Railway sets this)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default to 8080 locally
	}

	fmt.Println("UniGrade running on port:", port)
	http.ListenAndServe(":"+port, nil)
}
