package goo_oceanengine

import (
	"encoding/json"
	goo_http_request "github.com/liqiongtao/googo.io/goo-http-request"
)

type oceanengine struct {
	config Config
}

func (oe oceanengine) request(method, url string, data []byte, rst interface{}, opts ...goo_http_request.Option) (err error) {
	r := goo_http_request.New(opts...)

	if oe.config.Debug {
		r.Debug()
		r.SetHeader("X-Debug-Mode", "1")
	}

	var (
		b []byte
	)

	switch method {
	case "GET":
		r.SetContentType(goo_http_request.CONTENT_TYPE_JSON)
		if l := len(data); l > 0 {
			b, err = r.GetWithQuery(url, data)
		} else {
			b, err = r.Get(url)
		}
	case "POST":
		b, err = r.PostJson(url, data)
	}

	if err != nil {
		return
	}
	if err := json.Unmarshal(b, &rst); err != nil {
		return
	}

	return
}

func (oe oceanengine) get(url string, data []byte, rst interface{}, opts ...goo_http_request.Option) error {
	return oe.request("GET", url, data, rst, opts...)
}

func (oe oceanengine) post(url string, data []byte, rst interface{}, opts ...goo_http_request.Option) error {
	return oe.request("POST", url, data, rst, opts...)
}

func (oe oceanengine) Debug() oceanengine {
	oe.config.Debug = true
	return oe
}
