// Package raindrops contains convert function.
package raindrops

import "strconv"

// Convert converts numbers to strings.
func Convert(number int) string {
	raindrop := ""
	if number%3 == 0 || number%5 == 0 || number%7 == 0 {
		if number%3 == 0 {
			raindrop += "Pling"
		}
		if number%5 == 0 {
			raindrop += "Plang"
		}
		if number%7 == 0 {
			raindrop += "Plong"
		}
		return raindrop
	}
	return strconv.Itoa(number)

}
