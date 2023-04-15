package httputil

import (
	"github.com/SuCrayon/crayontool-go/pkg/constant"
	"net/http"
)

func AddCookie(req *http.Request, cookie *http.Cookie) {
	req.AddCookie(cookie)
}

func DelCookie(req *http.Request, name string) {
	cookie, _ := req.Cookie(name)
	cookie.MaxAge = -1
}

func GetCookieValue(req *http.Request, name string) (string, bool) {
	cookie, err := req.Cookie(name)
	if err != nil {
		return "", constant.False
	}
	return cookie.Value, constant.True
}
