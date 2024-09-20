package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"groupie/Errors"
	"groupie/Fetch"
)

func server(w http.ResponseWriter, r *http.Request) {
	// Use a switch statement to handle different URL paths
	jsonArtistsCards := Fetch.Fetch_cards(w, r)

	switch r.URL.Path {
	case "/":
		// Construct the path to the index.html template file
		IndexPath := filepath.Join("..", "templates", "index.html")

		// Parse the index.html template
		IndexParse, err := template.ParseFiles(IndexPath)
		if err != nil {
			// Call a custom error handler function for HTTP 500 errors
			Errors.Error500(w, r)
		}

		// Execute the index.html template
		err = IndexParse.ExecuteTemplate(w, "index.html", jsonArtistsCards)
		if err != nil {
			Errors.Error500(w, r)
			return
		}
	case "/profile":
		Profile_Path := filepath.Join("..", "templates", "profile.html")
		// Parse the index.html template
		Profile_Parse, err := template.ParseFiles(Profile_Path)

		//store all artist infos
		var artist_info interface{}

		var massage string

		artist_info, massage = Fetch.Fetch_profile(w, r, jsonArtistsCards)

		if massage == "500" {
			Errors.Error500(w, r)
		} else {
			Profile_Parse.ExecuteTemplate(w, "profile.html", artist_info)
			if err != nil {
				// Call a custom error handler function for HTTP 500 errors
				Errors.Error500(w, r)
			}
		}

	default:
		// Handle requests for paths other than "/"
		Errors.Error404(w, r)
	}
}

func main() {
	// Handle requests to the root URL ("/")
	http.HandleFunc("/", server)
	styles := http.FileServer(http.Dir("../stylesheets"))
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", styles))
	// Start the HTTP server on port 8080

	fmt.Println("Server up and runing on port :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server is down", err)
	}

}
