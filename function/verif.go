package groupie

var data ArtistStruct
func SubApiVerif(apifile string, number string) ArtistStruct {
	path := ""
	switch apifile {
	case "Locations":
		path = "https://groupietrackers.herokuapp.com/api/locations/" + number 
		data = APIRequestLocations(path)
	case "ConcertDates":
		path = "https://groupietrackers.herokuapp.com/api/dates/" + number
		data = APIRequestDates(path)
	case "Relations":
		path = "https://groupietrackers.herokuapp.com/api/relation/" + number
	}
	return data
}