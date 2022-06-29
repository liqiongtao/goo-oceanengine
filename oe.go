package goo_oceanengine

import (
	"encoding/json"
	"errors"
	goo_http_request "github.com/liqiongtao/googo.io/goo-http-request"
	"io"
	"time"
)

type oceanengine struct {
	config Config
}

func (oe oceanengine) get(url string, data []byte, rst interface{}, opts ...goo_http_request.Option) (err error) {
	r := request{
		debug:   oe.config.Debug,
		timeout: oe.config.Timeout,
		method:  "GET",
		url:     url,
		data:    data,
		opts:    opts,
	}

	for i := 0; i < 5; i++ {
		if err = r.doHandle(rst); err != nil {
			time.Sleep(time.Duration(i+1) * time.Second)
			continue
		}
		break
	}

	return
}

func (oe oceanengine) post(url string, data []byte, rst interface{}, opts ...goo_http_request.Option) (err error) {
	r := request{
		debug:   oe.config.Debug,
		timeout: oe.config.Timeout,
		method:  "POST",
		url:     url,
		data:    data,
		opts:    opts,
	}

	for i := 0; i < 5; i++ {
		if err = r.doHandle(rst); err != nil {
			time.Sleep(time.Duration(i+1) * time.Second)
			continue
		}
		break
	}

	return
}

func (oe oceanengine) uploadFile(url, fileName string, fh io.Reader, data map[string]string, rst interface{}, opts ...goo_http_request.Option) (err error) {
	r := request{
		debug:       oe.config.Debug,
		timeout:     oe.config.Timeout,
		method:      "UPLOAD",
		url:         url,
		fileField:   "file",
		fileName:    fileName,
		fileData:    data,
		fileHandler: fh,
		opts:        opts,
	}

	for i := 0; i < 5; i++ {
		if err = r.doHandle(rst); err != nil {
			time.Sleep(time.Duration(i+1) * time.Second)
			continue
		}
		break
	}

	return
}

type request struct {
	debug bool

	method  string
	url     string
	timeout time.Duration

	data []byte

	fileField   string
	fileName    string
	fileData    map[string]string
	fileHandler io.Reader

	opts []goo_http_request.Option
}

func (re request) doHandle(rst interface{}) (err error) {
	r := goo_http_request.New(re.opts...)

	if re.timeout == 0 {
		re.timeout = 30 * time.Second
	}
	r.SetTimeout(re.timeout)

	if re.debug {
		r.Debug()
		r.SetHeader("X-Debug-Mode", "1")
	}

	var (
		b []byte
	)

	switch re.method {
	case "GET":
		r.SetContentType(goo_http_request.CONTENT_TYPE_JSON)
		if l := len(re.data); l > 0 {
			b, err = r.GetWithQuery(re.url, re.data)
		} else {
			b, err = r.Get(re.url)
		}

	case "POST":
		b, err = r.PostJson(re.url, re.data)

	case "UPLOAD":
		b, err = r.Upload(re.url, re.fileField, re.fileName, re.fileHandler, re.fileData)
	}

	if err != nil {
		return
	}

	var rsp map[string]interface{}
	if err = json.Unmarshal(b, &rsp); err != nil {
		return
	}
	if rsp["code"].(float64) > 0 {
		err = errors.New(rsp["message"].(string))
		return
	}

	if err = json.Unmarshal(b, &rst); err != nil {
		return
	}

	return
}
