package auth

import (
	"errors"
	"net/http"
	"strings"
)

// this function extracts api key from header to get user information
// the api key will look something like ' ApiKey {keyvalue} '

func Getapikey(headers http.Header) (string, error) {
	apival := headers.Get("Authorization")

	if apival == "" {
		return "", errors.New("empty authenticaiton found no api key found on header")
	}
	vals := strings.Split(apival, " ")
	if len(vals) != 2 {
		return "", errors.New("malfomed authentication header")

	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malfomed authentication header first half")
	}
	return vals[1], nil
}
