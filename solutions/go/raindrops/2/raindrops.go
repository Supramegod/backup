package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	
	vowels := map[int]string{
		3: "i",
		5: "a", 
		7: "o",
	}


	factors := []int{3, 5, 7}

	var sounds []string

	for _, factor := range factors {
		if number%factor == 0 {

			sound := "Pl" + vowels[factor] + "ng"
			sounds = append(sounds, sound)
		}
	}

	if len(sounds) == 0 {
		return strconv.Itoa(number)
	}

	return strings.Join(sounds, "")
}