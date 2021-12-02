package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	requestURL = "https://api.wheretheiss.at/v1/satellites/25544"
	finalURL   = "https://www.openstreetmap.org/#map=12/"
)

type Location struct { //I only declare the values I need form the response
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func main() {
	response, err := http.Get(requestURL)
	if err != nil {
		log.Fatal("The API Call failed! ", err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Reading the data stream failed! ", err)
	}
	defer response.Body.Close()
	var locationData Location
	err = json.Unmarshal(data, &locationData)
	if err != nil {
		log.Fatal("Parsing the JSON data to a variable failed! ", err)
	}

	lat := fmt.Sprintf("%f", locationData.Latitude) //converting float64 to strings
	lon := fmt.Sprintf("%f", locationData.Longitude)
	volleOpenURL := finalURL + lat + "/" + lon
	log.Println(volleOpenURL) //Print it to a OpenMaps URL
}
