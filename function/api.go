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


var TabA []Artist
var TabL Locations
var TabD Dates

func APIRequestArtist() ArtistStruct {
	
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



func APIRequestLocations(apiloc string) ArtistStruct {
	req, err := http.Get(apiloc)
	if err != nil {
		fmt.Println(err)
	}
	d, err2 := ioutil.ReadAll(req.Body)
	if err2 != nil {
		fmt.Println(err2)
	}
	json.Unmarshal(d, &TabL)
	data := ArtistStruct {
		TabLocations: TabL,
	}
	return data
}

func APIRequestDates(apidates string) ArtistStruct {
	req, err := http.Get(apidates)
	if err != nil {
		fmt.Println(err)
	}
	d, err2 := ioutil.ReadAll(req.Body)
	if err2 != nil {
		fmt.Println(err2)
	}
	json.Unmarshal(d, &TabD)
	fmt.Println(TabD)
	data := ArtistStruct {
		TabDates: TabD,
	}
	return data
}