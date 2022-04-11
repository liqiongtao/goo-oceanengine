package goo_oceanengine

import (
	goo_http_request "github.com/liqiongtao/googo.io/goo-http-request"
)

func Campaign(config Config) campaign {
	return campaign{oceanengine{config: config}}
}

type campaign struct {
	oceanengine
}

// 获取广告组
// https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710532657164#item-link-%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80
// 目前广告组信息主要包括营销目的、推广目的、预算信息以及其他基础信息，此接口用于获取广告组列表的信息；
func (ca campaign) Get(params CampaignGetParams, accessToken string) (rst CampaignGetResult) {
	opt := goo_http_request.HeaderOption("Access-Token", accessToken)

	rst = CampaignGetResult{}
	if err := ca.get(CAMPAIGN_GET_URL, params.Json(), &rst, opt); err != nil {
		rst = CampaignGetResult{Code: 5001, Message: err.Error()}
	}
	return
}
