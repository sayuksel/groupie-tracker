package Fetch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"groupie/Errors"
)

type relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func Fetch_Relation(w http.ResponseWriter, r *http.Request, url string) [][]string {
	respose, err := http.Get(url)

	defer respose.Body.Close()

	data, err := ioutil.ReadAll(respose.Body)

	var jsonRelations relations

	err = json.Unmarshal(data, &jsonRelations)

	var dates_loca [][]string
	if err != nil {
		Errors.Error500(w, r)
	}
	// passing the data from the map to the dates_loca slice
	for location, dates := range jsonRelations.DatesLocations {
		locationAndDates := append(strings.Split(strings.ToUpper(location), "-"), dates...)
		dates_loca = append(dates_loca, locationAndDates)

	}

	return dates_loca
}
