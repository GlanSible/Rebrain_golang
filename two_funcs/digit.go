package two_funcs

var digits map[string]string = map[string]string{
	"One":	 "one",
	"Two":	 "two",
	"Three": "three",
	"Four":	 "four",
	"Five":	 "five",
}

func Digit(randomKey string) string {
	chosenDigit := digits[randomKey]
	return chosenDigit
}
