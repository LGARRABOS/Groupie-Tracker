func Filter(loc string, first int, create int, members int) ArtistStruct {
	for i = 1; i <= 52; i++ {
		if locVerif(i, loc) && firstVerif(i, first) && createVerif(i, create) {

		}
	}
}

func locVerif(i int, loc string) bool {
	path := "https://groupietrackers.herokuapp.com/api/locations/" + strconv.Itoa(i)
	data := APIRequest(path, "Locations")
	for i = 0; i < len(data.TabLocations.Locations); i++ {
		if data.TabLocations.Locations[i] == loc {
			return true
		}	
	}
	return false
}

func firstVerif(i int, first int) bool {
	path := "https://groupietrackers.herokuapp.com/api/artist/" + strconv.Itoa(i)
	data := APIRequest(path, "Artists")
	if data.TabArtist.FirstAlbum == strconv.Itoa(first) {
		return true
	}
	return false
}
func createVerif(i int, create int) bool {
	path := "https://groupietrackers.herokuapp.com/api/artist/" + strconv.Itoa(i)
	data := APIRequest(path, "Artists")
	date := ""
	for i = 0; i <= len(data.TabArtist.CreationDate); i++ {
		if i < 6 {
			date = date + data.TabArtist.CreationDate[i]
		}
	}
	if date == strconv.Itoa(first) {
		return true
	}
	return false
}