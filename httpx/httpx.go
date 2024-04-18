package httpx

import (
	"crypto/tls"
	"fmt"
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
		// 加载证书和私钥文件
		cert, err := tls.LoadX509KeyPair("storage/proxy/proxy.crt", "storage/proxy/proxy.key")
		if err != nil {
			fmt.Println("Error loading key pair:", err)
			return
		}
		httpClient = &http.Client{
			Timeout: time.Second * time.Duration(60),
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout: 5 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout: 5 * time.Second,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
					Certificates:       []tls.Certificate{cert},
				},
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 100,
				IdleConnTimeout:     90 * time.Second,
			},
		}
	})
}
