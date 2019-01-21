package space

// import (
// 	"fmt"
// 	"math"
// )
//
// func main() {
// 	// takes age in number of seconds and calculates how old
// 	// someone would be on each planet.
// 	var age int
// 	fmt.Println("How old are you in seconds? ")
// 	fmt.Scan(&age)
// 	float_input := float64(age)
// 	MercuryYear(float_input)
// 	VenusYear(float_input)
// 	EarthYear(float_input)
// 	MarsYear(float_input)
// 	JupiterYear(float_input)
// 	SaturnYear(float_input)
// 	UranusYear(float_input)
// 	NeptuneYear(float_input)
// }
//
// func MercuryYear(seconds float64) {
// 	fmt.Println(seconds)
// 	mercury_seconds := 0.2408467 * 31557600
// 	age := seconds / mercury_seconds
// 	fmt.Println("On Mercury you would be", math.Round(age*100)/100, "years old.")
// }
//
// func VenusYear(seconds float64) {
// 	venus_seconds := 0.61519726 * 31557600
// 	age := seconds / venus_seconds
// 	fmt.Println("On Venus you would be", math.Round(age*100)/100, "years old.")
// }
//
// func EarthYear(seconds float64) {
// 	age := seconds / 31557600
// 	fmt.Println("On Earth you are", math.Round(age*100)/100, "years old.")
// }
//
// func MarsYear(seconds float64) {
// 	mars_seconds := 1.8808158 * 31557600
// 	age := seconds / mars_seconds
// 	fmt.Println("On Mars you would be", math.Round(age), "years old.")
// }
//
// func JupiterYear(seconds float64) {
// 	jupiter_seconds := 11.862615 * 31557600
// 	age := seconds / jupiter_seconds
// 	fmt.Println("On Jupiter you would be", math.Round(age*100)/100, "years old.")
// }
//
// func SaturnYear(seconds float64) {
// 	saturn_seconds := 29.447498 * 31557600
// 	age := seconds / saturn_seconds
// 	fmt.Println("On Saturn you would be", math.Round(age*100)/100, "years old.")
// }
//
// func UranusYear(seconds float64) {
// 	uranus_seconds := 84.016846 * 31557600
// 	age := seconds / uranus_seconds
// 	fmt.Println("On Uranus you would be", math.Round(age*100)/100, "years old.")
// }
//
// func NeptuneYear(seconds float64) {
// 	neptune_seconds := 164.79132 * 31557600
// 	age := seconds / neptune_seconds
// 	fmt.Println("On Neptune you would be", math.Round(age*100)/100, "years old.")
// }
