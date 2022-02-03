package goo_oceanengine

import (
	goo_http_request "github.com/liqiongtao/googo.io/goo-http-request"
	goo_utils "github.com/liqiongtao/googo.io/goo-utils"
)

func AD(config Config) ad {
	return ad{oceanengine{config: config}}
}

type ad struct {
	oceanengine
}

// 获取广告计划
// https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710535150604#item-link-%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80
// 此接口用于获取广告计划列表的信息；
func (ad_ ad) Get(params AdGetParams, accessToken string) goo_utils.Params {
	return ad_.get(AD_GET_URL, params.Json(), goo_http_request.HeaderOption("Access-Token", accessToken))
}
