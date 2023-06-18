package utills

import (
	"strings"
	)

//BuildFullName will create a full name based on first and last names
func BuildFullName(firstname, lastname string) string {
	return strings.ToLower(firstname) + " " + strings.ToLower(lastname)
}