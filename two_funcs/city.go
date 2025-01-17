package two_funcs

var cities map[string]string = map[string]string{
	"One":	 "Moscow",
	"Two":	 "Tokyo",
	"Three": "San-Francisco",
	"Four":	 "Boston",
	"Five":	 "Lissboa",
}

func City(randomKey string) string {
	chosenCity := cities[randomKey]
	return	chosenCity
}
