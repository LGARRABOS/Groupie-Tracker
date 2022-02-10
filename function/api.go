package groupie

import  (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


type ArtistStruct struct {
	TabArtist []Artist
	TabLocations Locations
	TabDates Dates
}

type Artist struct {
	Id	int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Member []string`json:"members"`
	CreationDate int`json:"creationDate"`
	FirstAlbum string `json:"firstAlbum"`
	Locations string `json:"locations"`
	ConcertDates string `json:"concertDates"`
	Relations string `json:"relations"`
}

type Locations struct {
	Id	int `json:"id"`
	Locations []string `json:"locations"`
	Dates string `json:"dates"`	
}
type Dates struct {
	Id	int `json:"id"`
	Dates []string `json:"dates"`
}






func APIRequestArtist() ArtistStruct {
	var TabA []Artist
	req, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println(err)
	}
	d, err2 := ioutil.ReadAll(req.Body)
	if err2 != nil {
		fmt.Println(err2)
	}
	json.Unmarshal(d, &TabA)
	data := ArtistStruct {
		TabArtist: TabA,
	}
	return data
}



func APIRequest(apiloc string, who string) ArtistStruct {
	req, err := http.Get(apiloc)
	if err != nil {
		fmt.Println(err)
	}
	d, err2 := ioutil.ReadAll(req.Body)
	if err2 != nil {
		fmt.Println(err2)
	}
	switch who {
	case "Locations" :
		var TabL Locations
		json.Unmarshal(d, &TabL)
		data := ArtistStruct {
			TabLocations: TabL,
		}
		return data
	case "ConcertDates":
		var TabD Dates
		json.Unmarshal(d, &TabD)
		data := ArtistStruct {
			TabDates: TabD,
		}
		return data
	case "Relations":

	}
	return data
}

