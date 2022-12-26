package httputil

import "net/http"

const (
	HeaderContentType = "Content-Type"
	HeaderAuthorization = "Authorization"
)

const (
	ApplicationJSON = "application/json"
)

func SetContentType(req *http.Request, value string) {
	req.Header.Set(HeaderContentType, value)
}

func AddContentType(req *http.Request, value string) {
	req.Header.Add(HeaderContentType, value)
}

func SetApplicationJSONHeader(req *http.Request) {
	SetContentType(req, ApplicationJSON)
}

func AddApplicationJSONHeader(req *http.Request) {
	AddContentType(req, ApplicationJSON)
}

func SetAuthorization(req *http.Request, value string) {
	req.Header.Set(HeaderAuthorization, value)
}

func AddAuthorization(req *http.Request, value string) {
	req.Header.Add(HeaderAuthorization, value)
}

func SetBasicAuthHeader(req *http.Request, username, password string) {
	req.SetBasicAuth(username, password)
}

func SetHeaderFromMap(req *http.Request, header map[string]string) {
	for k, v := range header {
		req.Header.Set(k, v)
	}
}

func AddHeaderFromMap(req *http.Request, header map[string]string) {
	for k, v := range header {
		req.Header.Add(k, v)
	}
}

func DelHeaderFromMap(req *http.Request, header map[string]string) {
	for k := range header {
		req.Header.Del(k)
	}
}

func DelHeaderFromSlice(req *http.Request, header []string) {
	for _, k := range header {
		req.Header.Del(k)
	}
}

func DelHeader(req *http.Request, header ...string) {
	for _, k := range header {
		req.Header.Del(k)
	}
}