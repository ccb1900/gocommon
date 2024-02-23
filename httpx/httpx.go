package httpx

import (
	"crypto/tls"
	"net"
	"net/http"
	"sync"
	"time"
)

var (
	once       sync.Once
	httpClient *http.Client
)

func Default() *http.Client {
	return httpClient
}

func Init() {
	once.Do(func() {
		httpClient = &http.Client{
			Timeout: time.Second * time.Duration(60),
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout: 5 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout: 5 * time.Second,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 100,
				IdleConnTimeout:     90 * time.Second,
			},
		}
	})
}
