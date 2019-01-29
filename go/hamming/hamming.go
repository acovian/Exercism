package hamming

import "errors"

// Distance compares two strings and determines which instances are
// unequal.
func Distance(a, b string) (int, error) {
	counter := 0
	if len(b) == len(a) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				counter++
			}
		}
		return counter, nil
	}
	return counter, errors.New("strings are of unequal length")
}
