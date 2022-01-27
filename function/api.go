package groupie

import  (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ArtistStruct struct {
	Tab []Artist
}

type Artist struct {
	Id	int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Menber []string`json:"members"`
	CreationDate int`json:"creationDate"`
	FrstAlbum string `json:"firstAlbum"`
}

var ArtistTab []Artist

func APIRequest() ArtistStruct {
	req, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		fmt.Println(err)
	}

	d, err2 := ioutil.ReadAll(req.Body)

	if err2 != nil {
		fmt.Println(err2)
	}
	json.Unmarshal(d, &ArtistTab)
	data := ArtistStruct {
		Tab: ArtistTab,
	}
	return data

}
