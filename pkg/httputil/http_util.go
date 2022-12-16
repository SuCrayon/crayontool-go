package httputil

import (
	"net/http"
	_httputil "net/http/httputil"
	"net/url"
)

func NewSimpleForwardProxy(forwardTo string) (*_httputil.ReverseProxy, error) {
	u, err := url.Parse(forwardTo)
	if err != nil {
		return nil, err
	}
	return &_httputil.ReverseProxy{
		Director: func(request *http.Request) {
			request.URL = u
		},
	}, nil
}
