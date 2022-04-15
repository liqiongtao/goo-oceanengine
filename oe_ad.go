package goo_oceanengine

import (
	goo_http_request "github.com/liqiongtao/googo.io/goo-http-request"
)

func AD(config Config) ad {
	return ad{oceanengine{config: config}}
}

type ad struct {
	oceanengine
}

// 获取广告计划
// https://open.oceanengine.com/labels/7/docs/1696710535150604
// 此接口用于获取广告计划列表的信息；
func (ad_ ad) Get(params AdGetParams, accessToken string) (rst AdGetResult) {
	opt := goo_http_request.HeaderOption("Access-Token", accessToken)

	rst = AdGetResult{}
	if err := ad_.get(AD_GET_URL, params.Json(), &rst, opt); err != nil {
		rst = AdGetResult{Code: 5001, Message: err.Error()}
	}
	return
}
