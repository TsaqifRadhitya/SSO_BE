package Utils

import (
	"errors"
	"strings"
)

func ExtractBearerToken(BearerToken string) (string, error) {
	if BearerToken == "" {
		return "", errors.New("BearerToken is empty")
	}

	parts := strings.SplitN(BearerToken, " ", 2)

	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("Invalid BearerToken")
	}

	return parts[1], nil
}
