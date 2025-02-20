package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)

	if err != nil {
		return "", err
	}

	normalizedURL := parsedURL.Host + parsedURL.Path
	return strings.TrimSuffix(normalizedURL, "/"), nil
}
