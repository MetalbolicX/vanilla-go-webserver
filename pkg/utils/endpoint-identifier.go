package utils

import "strings"

// GetIdentifier extracts the last part of the endPoint and returns it.
func GetIdentifier(endPoint string) string {
	endPointSegments := strings.Split(endPoint, "/")
	return endPointSegments[len(endPointSegments)-1]
}
