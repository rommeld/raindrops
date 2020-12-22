package raindrops

import (
	"strconv"
)

//Convert a integer based on the moduli into a string of rain drops
func Convert(input int) (rainSounds string) {
	if input%3 == 0 {
		rainSounds += "Pling"
	}
	if input%5 == 0 {
		rainSounds += "Plang"
	}
	if input%7 == 0 {
		rainSounds += "Plong"
	}
	if len(rainSounds) == 0 {
		return strconv.Itoa(input)
	}
	return rainSounds
}
