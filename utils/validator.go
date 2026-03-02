package utils

import "strings"

func IsKIITEmail(email string) bool {
	return strings.HasSuffix(email, "@kiit.ac.in")
}
