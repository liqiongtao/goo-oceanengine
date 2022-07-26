package goo_oceanengine

import (
	"encoding/json"
)

// 获取广告组 - 请求参数
type CampaignGetParams struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
	filtering    struct {
		Ids                []int64 `json:"ids,omitempty"`                  // 广告组ID过滤，数组，不超过100个
		CampaignName       string  `json:"campaign_name,omitempty"`        // 广告组name过滤，长度为1-30个字符，其中1个中文字符算2位
		LandingType        string  `json:"landing_type,omitempty"`         // 广告组推广目的过滤，详见【附录-推广目的类型】
		Status             string  `json:"status,omitempty"`               // 广告组状态过滤，默认为返回“所有不包含已删除”，如果要返回所有包含已删除有对应枚举表示，详见【附录-广告组状态】
		CampaignCreateTime string  `json:"campaign_create_time,omitempty"` // 广告组创建时间，格式yyyy-mm-dd，表示过滤出当天创建的广告组
	} `json:"filtering,omitempty"` // 过滤条件，若此字段不传，或传空则视为无限制条件
	Fields   []string `json:"fields,omitempty"`    // 查询字段集合，如果指定，则返回结果数组中， 每个元素是包含所查询字段的字典；允许值:id、name、budget、 budget_mode、landing_type、status、modify_time、status、modify_time、campaign_modify_time、campaign_create_time
	Page     int32    `json:"page,omitempty"`      // 页数默认值: 1
	PageSize int32    `json:"page_size,omitempty"` // 页面大小默认值:10，page_size范围为[1,1000]
}

func (p CampaignGetParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type CampaignGetResult struct {
	Code      int    `json:"code" bson:"code"`
	Message   string `json:"message" bson:"message"`
	RequestId string `json:"request_id" bson:"request_id"`
	Data      struct {
		List     []CampaignInfo `json:"list" bson:"list"`
		PageInfo struct {
			Page        int `json:"page" bson:"page"`
			PageSize    int `json:"page_size" bson:"page_size"`
			TotalNumber int `json:"total_number" bson:"total_number"`
			TotalPage   int `json:"total_page" bson:"total_page"`
		} `json:"page_info" bson:"page_info"`
	} `json:"data" bson:"data"`
}

func (p CampaignGetResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type CampaignInfo struct {
	MarketingScene       string  `json:"marketing_scene" bson:"marketing_scene"`
	DedicateType         string  `json:"dedicate_type" bson:"dedicate_type"`
	ModifyTime           string  `json:"modify_time" bson:"modify_time"`
	Name                 string  `json:"name" bson:"name"`
	SmartBidType         string  `json:"smart_bid_type" bson:"smart_bid_type"`
	DeliveryMode         string  `json:"delivery_mode" bson:"delivery_mode"`
	CampaignId           int64   `json:"campaign_id" bson:"campaign_id"`
	Budget               float64 `json:"budget" bson:"budget"`
	CampaignCreateTime   string  `json:"campaign_create_time" bson:"campaign_create_time"`
	CampaignModifyTime   string  `json:"campaign_modify_time" bson:"campaign_modify_time"`
	LandingType          string  `json:"landing_type" bson:"landing_type"`
	MarketingPurpose     string  `json:"marketing_purpose" bson:"marketing_purpose"`
	Id                   int64   `json:"id" bson:"id"`
	AdvertiserId         int64   `json:"advertiser_id" bson:"advertiser_id"`
	CampaignType         string  `json:"campaign_type" bson:"campaign_type"`
	Status               string  `json:"status" bson:"status"`
	BudgetMode           string  `json:"budget_mode" bson:"budget_mode"`
	DeliveryRelatedNum   string  `json:"delivery_related_num" bson:"delivery_related_num"`
	BudgetOptimizeSwitch string  `json:"budget_optimize_switch" bson:"budget_optimize_switch"`
}
