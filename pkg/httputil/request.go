package httputil

import (
	"bytes"
	"crayontool-go/pkg/constant"
	"crayontool-go/pkg/logger"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	JSONMarshalErr = errors.New("json marshal failed")
	NewRequestErr = errors.New("new http request failed")
)

var (
	defaultTimeout = 20 * time.Second
)

func NewTLSInsecureClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: constant.True,
			},
		},
		Timeout: defaultTimeout,
	}
}

func NewJSONRequest(method, url string, body interface{}) (*http.Request, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		logger.Alert(fmt.Sprintf("%s, err: %+v\n", JSONMarshalErr.Error(), err))
		return nil, err
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(jsonBody))
	if err != nil {
		logger.Alert(fmt.Sprintf("%s, err: %+v\n", NewRequestErr.Error(), err))
		return nil, err
	}
	// 设置请求头
	AddApplicationJSONHeader(req)
	return req, nil
}