package cmd

import (
	"net/url"
	"strings"
)

func extractDomain(urlString string) (string, error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}

	domain := strings.TrimPrefix(parsedURL.Host, "www.")
	return domain, nil
}

func removeProtocol(urlString string) (string, error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}

	return parsedURL.Host + parsedURL.Path + parsedURL.RawQuery + "#" + parsedURL.Fragment, nil
}

func removeDomain(urlString string, domain string) string {
	prefix := domain + "/"
	return strings.TrimPrefix(urlString, prefix)
}
