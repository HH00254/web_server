package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts an API Key from
// the headrers of an HTTP request
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	value := headers.Get("Authorization")
	if value == "" {
		return "", errors.New("No Authorization header provided")
	}

	vals := strings.Split(value, " ")
	if len(vals) != 2 {
		return "", errors.New("Malformed Authorization header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed Authorization header")
	}

	return vals[1], nil
}
