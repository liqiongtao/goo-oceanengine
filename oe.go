package goo_oceanengine

import (
	"encoding/json"
	goo_http_request "github.com/liqiongtao/googo.io/goo-http-request"
	goo_log "github.com/liqiongtao/googo.io/goo-log"
	"time"
)

type oceanengine struct {
	config Config
}

func (oe oceanengine) request(method, url string, data []byte, rst interface{}, opts ...goo_http_request.Option) (err error) {
	r := goo_http_request.New(opts...)

	r.SetTimeout(30 * time.Second)

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
	if err = json.Unmarshal(b, &rst); err != nil {
		return
	}

	return
}

func (oe oceanengine) get(url string, data []byte, rst interface{}, opts ...goo_http_request.Option) (err error) {
	for i := 0; i < 5; i++ {
		if err = oe.request("GET", url, data, rst, opts...); err != nil {
			time.Sleep(time.Duration(i+1) * time.Second)
			goo_log.ErrorF("重试第%d次, %s, %s", i+1, url, string(data))
			continue
		}
		break
	}
	return
}

func (oe oceanengine) post(url string, data []byte, rst interface{}, opts ...goo_http_request.Option) (err error) {
	for i := 0; i < 5; i++ {
		if err = oe.request("POST", url, data, rst, opts...); err != nil {
			time.Sleep(time.Duration(i+1) * time.Second)
			goo_log.ErrorF("重试第%d次, %s, %s", i+1, url, string(data))
			continue
		}
		break
	}
	return
}

func (oe oceanengine) Debug() oceanengine {
	oe.config.Debug = true
	return oe
}
