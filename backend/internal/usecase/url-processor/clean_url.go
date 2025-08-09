package urlprocessor

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

func (u *usecase) CleanURL(ctx context.Context, operation, uri string) (string, error) {
	parsed, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	switch strings.ToLower(operation) {
	case "canonical":
		applyCanonical(parsed)
	case "redirection":
		applyRedirection(parsed)
	case "all":
		applyCanonical(parsed)
		applyRedirection(parsed)
	default:
		return "", fmt.Errorf("invalid operation")
	}

	return parsed.String(), nil
}

func applyCanonical(u *url.URL) {
	// remove query parameters
	u.RawQuery = ""
	// remove trailing slash from the path
	u.Path = strings.TrimSuffix(u.Path, "/")
}

func applyRedirection(u *url.URL) {
	// force domain redirect to www.byfood.com
	u.Host = "www.byfood.com"
	u.Scheme = "https"
	// lowercase entire URL parts
	u.Path = strings.ToLower(u.Path)
}
