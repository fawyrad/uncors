package infra

import (
	"net/http"
	"time"

	"github.com/evg4b/uncors/internal/sfmt"
	"github.com/evg4b/uncors/pkg/urlx"
)

const defaultTimeout = 5 * time.Minute

var defaultHTTPClient = http.Client{
	CheckRedirect: func(r *http.Request, v []*http.Request) error {
		return http.ErrUseLastResponse
	},
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	},
	Jar:     nil,
	Timeout: defaultTimeout,
}

func MakeHTTPClient(proxy string) (*http.Client, error) {
	if len(proxy) > 0 {
		parsedURL, err := urlx.Parse(proxy)
		if err != nil {
			return nil, sfmt.Errorf("failed to create http client: %w", err)
		}

		httpClient := defaultHTTPClient
		httpClient.Transport = &http.Transport{
			Proxy: http.ProxyURL(parsedURL),
		}

		return &httpClient, nil
	}

	return &defaultHTTPClient, nil
}
