package httputil

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

func BodyReadAll(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func UnmarshalJSONResp(resp *http.Response, v interface{}) error {
	bs, err := BodyReadAll(resp)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, v)
}

func UnmarshalXMLResp(resp *http.Response, v interface{}) error {
	bs, err := BodyReadAll(resp)
	if err != nil {
		return err
	}
	return xml.Unmarshal(bs, v)
}
