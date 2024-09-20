package Fetch

import (
	"net/http"
	"strconv"
)

func Fetch_profile(w http.ResponseWriter, r *http.Request, jsonArtistsCards []Artist) (*Artist, string) {

	var artistInfo *Artist

	// Parse artist ID from URL query parameter
	artistID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return artistInfo, "500"
	}
	var data_location [][]string
	if artistID < 1 || artistID > len(jsonArtistsCards) {
		return artistInfo, "500"
	}

	index := artistID
	index--
	data_location = Fetch_Relation(w, r, jsonArtistsCards[index].Relations)

	// Find the artist in jsonArtistsCards
	for _, artist := range jsonArtistsCards {
		if artist.ID == artistID {
			artistInfo = &artist
			break
		}
	}
	artistInfo.Date_Locat = append(artistInfo.Date_Locat, data_location...)

	// Handle case where artist is not found
	if artistInfo == nil {
		return artistInfo, "500"
	}

	// Return the artist information
	return artistInfo, ""
}
