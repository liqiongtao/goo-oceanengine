package goo_oceanengine

import (
	"encoding/json"
)

// 获取广告计划 - 请求参数
type AdGetParams struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
	Filtering    struct {
		Ids          []int64  `json:"ids,omitempty"`            // 按广告计划ID过滤，范围为1-100
		AdName       string   `json:"ad_name,omitempty"`        // 按广告计划name过滤，长度为1-30个字符
		PricingList  []string `json:"pricing_list,omitempty"`   // 按出价方式过滤，详见【附录-计划出价类型】
		Status       string   `json:"status,omitempty"`         // 按计划状态过滤，默认为返回“所有不包含已删除”，如果要返回所有包含已删除有对应枚举表示，详见【附录-广告计划投放状态】注：暂不支持使用AD_STATUS_ADVERTISER_BUDGET_EXCEED筛选
		CampaignId   int64    `json:"campaign_id,omitempty"`    // 按广告组id过滤
		AdCreateTime string   `json:"ad_create_time,omitempty"` // 广告计划创建时间，格式yyyy-mm-dd，表示过滤出当天创建的广告计划
		AdModifyTime string   `json:"ad_modify_time,omitempty"` // 广告计划更新时间，格式yyyy-mm-dd，表示过滤出当天更新的广告计划
	} `json:"filtering,omitempty"` // 过滤条件，若此字段不传，或传空则视为无限制条件
	Fields   []string `json:"fields,omitempty"`    // 查询字段集合, 如果指定, 则返回结果数组中，不支持audience下字段筛选
	Page     int32    `json:"page,omitempty"`      // 页数默认值: 1，page范围为[1,99999]
	PageSize int32    `json:"page_size,omitempty"` // 页面大小默认值:10，page_size范围为[1,100]
}

func (p AdGetParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type AdGetResult struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		List []struct {
			CampaignId   int    `json:"campaign_id"`
			Id           int    `json:"id"`
			Name         string `json:"name"`
			AdCreateTime string `json:"ad_create_time"`
			AdModifyTime string `json:"ad_modify_time"`
		} `json:"list"`
		PageInfo struct {
			Page        int `json:"page"`
			PageSize    int `json:"page_size"`
			TotalNumber int `json:"total_number"`
			TotalPage   int `json:"total_page"`
		} `json:"page_info"`
	} `json:"data"`
}
