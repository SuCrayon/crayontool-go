package httputil

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNewServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		name := request.FormValue("name")
		writer.Write([]byte(fmt.Sprintf("your name is %s", name)))
	})
	mux.HandleFunc("/proxy", func(writer http.ResponseWriter, request *http.Request) {
		to := request.FormValue("to")
		proxy, err := NewSimpleForwardProxy(to)
		if err != nil {
			return
		}
		proxy.ServeHTTP(writer, request)
	})
	err := http.ListenAndServe(":28080", mux)
	if err != nil {
		panic(err)
	}
}
