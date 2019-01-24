// Package bob differentiates various inputs to output specific responses.
package bob

import (
	"regexp"
	"strings"
)

// Hey provides responses to remarks.
func Hey(remark string) string {
	switch {
	case MatchNumber(remark) && strings.HasSuffix(remark, "?"):
		return "Sure."
	case MatchNumber(remark) && strings.Contains(remark, "a"):
		return "Whatever."
	case MatchNumber(remark) && remark == strings.ToUpper(remark):
		return "Whoa, chill out!"
	case remark == strings.ToUpper(remark) && strings.HasSuffix(remark, "?"):
		return "Calm down, I know what I'm doing!"
	case strings.HasSuffix(remark, "?"):
		return "Sure."
	case remark == strings.ToUpper(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func MatchNumber(s string) bool {
	matched, _ := regexp.MatchString(("[\\d]"), s)
	return matched
}
