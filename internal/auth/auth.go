package internal

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {

	api_key_value := headers.Get("ApiKey")

	if api_key_value == "" {
		return "", errors.New("sorry please provide api key")
	}
	headerValues := strings.Split(api_key_value, "")

	if len(headerValues) != 64 {
		return "", errors.New("invalid apikeys sent")
	}

	return api_key_value, nil
}
