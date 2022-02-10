package groupie

func SeparateString(api string) (string, string) {
	other := false
	apitype := ""
	number := ""
	for i:= 0; i < len(api); i++ {
		if api[i] == ' '  {
			other = true
		} else if other {
			number = number + string(api[i])
		} else {
			apitype = apitype + string(api[i])
		}
	}
	return number, apitype
}




