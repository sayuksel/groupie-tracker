package Fetch

import (
	"encoding/json"
	"io"
	"net/http"

	"groupie/Errors"
)

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	//Locations    string   `json:"locations"`
	//ConcertDates string   `json:"concertDates"`
	Relations string `json:"relations"`

	Date_Locat [][]string
}

var All_artists []Artist

func Fetch_cards(w http.ResponseWriter, r *http.Request) []Artist {
	// Define a function to fetch data from the artists API
	const artists = "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(artists)
	if err != nil {
		Errors.Error500(w, r)
		return nil
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		Errors.Error500(w, r)
		return nil
	}

	err = json.Unmarshal(data, &All_artists)
	if err != nil {
		Errors.Error500(w, r)
		return nil
	}
	return All_artists
}
