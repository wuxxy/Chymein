package Core

import (
	"context"
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/labstack/echo/v4"
)

func FrontendProxy(target string) echo.HandlerFunc {
	parsedURL, err := url.Parse(target)
	if err != nil {
		panic("invalid proxy target: " + err.Error())
	}

	proxy := httputil.NewSingleHostReverseProxy(parsedURL)

	// Fix outgoing host header
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = parsedURL.Host
		req.Header.Set("X-Forwarded-Host", req.Host)
		req.Header.Set("X-Forwarded-Proto", "http")
	}

	// Silences noisy dev reload cancellations
	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
		if errors.Is(err, context.Canceled) {
			return // suppress harmless noise
		}
		http.Error(rw, "Proxy error: "+err.Error(), http.StatusBadGateway)
	}

	return func(c echo.Context) error {
		proxy.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
