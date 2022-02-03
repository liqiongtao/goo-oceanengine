package goo_oceanengine

import (
	goo_http_request "github.com/liqiongtao/googo.io/goo-http-request"
	goo_log "github.com/liqiongtao/googo.io/goo-log"
	goo_utils "github.com/liqiongtao/googo.io/goo-utils"
)

type oceanengine struct {
	config Config
	debug  bool
}

func (oe oceanengine) request(method, url string, data []byte, opts ...goo_http_request.Option) (rst goo_utils.Params) {
	rst = goo_utils.NewParams()

	r := goo_http_request.New(opts...)
	log := goo_log.WithField("url", url).WithField("data", string(data))

	if oe.debug {
		r.SetHeader("X-Debug-Mode", "1")
	}

	var (
		b   []byte
		err error
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
		rst.Set("code", 5001).Set("message", err)
		log.Error(err)
		return
	}

	log.WithField("result", string(b))

	if rst, err = goo_utils.Byte(b).Params(); err != nil {
		rst.Set("code", 5002).Set("message", err)
		log.Error(err)
		return
	}

	if oe.debug {
		log.Debug()
	}

	return
}

func (oe oceanengine) get(url string, data []byte, opts ...goo_http_request.Option) (rst goo_utils.Params) {
	return oe.request("GET", url, data, opts...)
}

func (oe oceanengine) post(url string, data []byte, opts ...goo_http_request.Option) (rst goo_utils.Params) {
	return oe.request("POST", url, data, opts...)
}

func (oe oceanengine) Debug() oceanengine {
	oe.debug = true
	return oe
}
