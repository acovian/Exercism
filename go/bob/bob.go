// Taken from oneKelvinSmith's solution, sans Remark struct and newRemark
// function.
package bob

import (
	"strings"
	"unicode"
)

// Returns responses to remarks.
func Hey(remark string) string {
	r := strings.TrimSpace(remark)
	switch {
	case Silent(r):
		return "Fine. Be that way!"
	case Yell_Question(r):
		return "Calm down, I know what I'm doing!"
	case Yell(r):
		return "Whoa, chill out!"
	case Question(r):
		return "Sure."
	default:
		return "Whatever."
	}
}

func Silent(remark string) bool {
	return remark == ""
}

func Yell(remark string) bool {
	hasLetters := strings.IndexFunc(remark, unicode.IsLetter) >= 0
	isUppercased := strings.ToUpper(remark) == remark
	return hasLetters && isUppercased
}

func Question(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func Yell_Question(remark string) bool {
	return Yell(remark) && Question(remark)
}
