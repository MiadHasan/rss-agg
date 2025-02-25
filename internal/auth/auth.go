package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts api key from the request header
// Example: ApiKey {insert api key}
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	vals := strings.Split(val, " ")
	if len(vals) != 2 && vals[0] != "ApiKey" {
		return "", errors.New("malformed api key")
	}
	return vals[1], nil
}
