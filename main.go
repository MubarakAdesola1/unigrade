package main

import (
	"fmt"
	"net/http"
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

	fmt.Println("UniGrade running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
