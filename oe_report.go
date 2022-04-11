package goo_oceanengine

import (
	goo_http_request "github.com/liqiongtao/googo.io/goo-http-request"
)

func Report(config Config) report {
	return report{oceanengine{config: config}}
}

type report struct {
	oceanengine
}

// 广告主数据
// https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710550620160#item-link-%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80
// 此接口用于获取广告账户维度的投放数据，包括消耗、点击、展示等指标，具体可以参考应答参数指标说明。
// 数据更新频率:
// 数据5～10分钟更新一次
// 一般历史数据都不会变，除了数据有问题有校对的情况会更新历史数据，第二天10点可以获取前一天稳定的消耗数据
func (r report) AdvertiserGet(params ReportAdvertiserGetParams, accessToken string) (rst ReportAdvertiserGetResult) {
	opt := goo_http_request.HeaderOption("Access-Token", accessToken)

	rst = ReportAdvertiserGetResult{}
	if err := r.get(REPORT_ADVERTISER_GET_URL, params.Json(), &rst, opt); err != nil {
		rst = ReportAdvertiserGetResult{Code: 5001, Message: err.Error()}
	}
	return
}

// 广告组数据
// https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710551161856#item-link-%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80
// 此接口用于获取广告组纬度的投放数据，包括消耗、点击、展示等指标，具体可以参考应答参数指标说明。
// 数据更新频率:
// 数据5～10分钟更新一次
// 一般历史数据都不会变，除了数据有问题有校对的情况会更新历史数据，第二天10点可以获取前一天稳定的消耗数据
func (r report) CampaignGet(params ReportCampaignGetParams, accessToken string) (rst ReportCampaignGetResult) {
	opt := goo_http_request.HeaderOption("Access-Token", accessToken)

	rst = ReportCampaignGetResult{}
	if err := r.get(REPORT_CAMPAIGN_GET_URL, params.Json(), &rst, opt); err != nil {
		rst = ReportCampaignGetResult{Code: 5001, Message: err.Error()}
	}
	return
}

// 广告计划数据
// https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710551666703#item-link-%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80
// 此接口用于获取广告组纬度的投放数据，包括消耗、点击、展示等指标，具体可以参考应答参数指标说明。
// 数据更新频率:
// 数据5～10分钟更新一次
// 一般历史数据都不会变，除了数据有问题有校对的情况会更新历史数据，第二天10点可以获取前一天稳定的消耗数据
func (r report) AdGet(params ReportAdGetParams, accessToken string) (rst ReportAdGetResult) {
	opt := goo_http_request.HeaderOption("Access-Token", accessToken)

	rst = ReportAdGetResult{}
	if err := r.get(REPORT_AD_GET_URL, params.Json(), &rst, opt); err != nil {
		rst = ReportAdGetResult{Code: 5001, Message: err.Error()}
	}
	return
}

// 广告创意数据
// https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710552261644#item-link-%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80
// 此接口用于获取广告组纬度的投放数据，包括消耗、点击、展示等指标，具体可以参考应答参数指标说明。
// 数据更新频率:
// 数据5～10分钟更新一次
// 一般历史数据都不会变，除了数据有问题有校对的情况会更新历史数据，第二天10点可以获取前一天稳定的消耗数据
func (r report) CreativeGet(params ReportCreativeGetParams, accessToken string) (rst ReportCreativeGetResult) {
	opt := goo_http_request.HeaderOption("Access-Token", accessToken)
	rst = ReportCreativeGetResult{}
	if err := r.get(REPORT_CREATIVE_GET_URL, params.Json(), &rst, opt); err != nil {
		rst = ReportCreativeGetResult{Code: 5001, Message: err.Error()}
	}
	return
}
