package goo_oceanengine

import (
	"encoding/json"
)

// 广告主数据 - 请求参数
type ReportAdvertiserGetParams struct {
	AdvertiserId    string   `json:"advertiser_id,omitempty"`    // 广告主ID
	StartDate       string   `json:"start_date,omitempty"`       // 起始日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期
	EndDate         string   `json:"end_date,omitempty"`         // 结束日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期，时间跨度不能超过30天
	Fields          []string `json:"fields,omitempty"`           // 指定需要的指标名称，可参考应答参数返回的消耗指标字段 默认值：cost、show、avg_show_cost、click、ctr、avg_click_cost、convert、convert_rate、convert_cost
	GroupBy         []string `json:"group_by,omitempty"`         // 分组条件 默认为STAT_GROUP_BY_FIELD_STAT_TIME，支持多种分组条件，具体详见【分组组合规则】
	TimeGranularity string   `json:"time_granularity,omitempty"` // 时间粒度 默认值: STAT_TIME_GRANULARITY_DAILY 允许值:STAT_TIME_GRANULARITY_DAILY (按天维度),STAT_TIME_GRANULARITY_HOURLY (按小时维度)
	Filtering       struct {
		DeliveryMode []string `json:"delivery_mode,omitempty"` // 投放模式筛选。 允许值: STANDARD:标准投放 ADLAB_FREE:管家&省心投放
	} `json:"filtering,omitempty"`              // 过滤条件，若此字段不传，或传空则视为无限制条件
	Page     int32 `json:"page,omitempty"`      // 页数默认值: 1
	PageSize int32 `json:"page_size,omitempty"` // 页面大小默认值:20，page_size范围为[1,1000]
}

func (p ReportAdvertiserGetParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

// 广告组数据 - 请求参数
type ReportCampaignGetParams struct {
	AdvertiserId    string   `json:"advertiser_id,omitempty"`    // 广告主ID
	StartDate       string   `json:"start_date,omitempty"`       // 起始日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期
	EndDate         string   `json:"end_date,omitempty"`         // 结束日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期，时间跨度不能超过30天
	Fields          []string `json:"fields,omitempty"`           // 指定需要的指标名称，可参考应答参数返回的消耗指标字段 默认值：cost、show、avg_show_cost、click、ctr、avg_click_cost、convert、convert_rate、convert_cost
	GroupBy         []string `json:"group_by,omitempty"`         // 分组条件 默认为 STAT_GROUP_BY_FIELD_STAT_TIME。支持多种分组条件，具体详见【分组组合规则】
	TimeGranularity string   `json:"time_granularity,omitempty"` // 时间粒度 默认值: STAT_TIME_GRANULARITY_DAILY 允许值:STAT_TIME_GRANULARITY_DAILY (按天维度),STAT_TIME_GRANULARITY_HOURLY (按小时维度)
	Filtering       struct {
		CampaignIds  []int64  `json:"campaign_ids,omitempty"`  // 广告组id列表：按照campaign_id过滤，最多支持100个
		CampaignName string   `json:"campaign_name,omitempty"` // 广告组名称：按照campaign_name过滤，字符串长度限制1-30
		LandingTypes []string `json:"landing_types,omitempty"` // 推广目的列表：按照广告组推广目的过滤,详见【附录-推广目的类型】 允许值:"LINK","APP","DPA","GOODS","STORE","SHOP","AWEME"
		Status       string   `json:"status,omitempty"`        // 广告组状态：按照广告组状态过滤，默认为返回“所有不包含已删除”，如果要返回所有包含已删除有对应枚举表示，详见【附录-广告组状态】
		DeliveryMode []string `json:"delivery_mode,omitempty"` // 投放模式筛选。 允许值: STANDARD:标准投放 ADLAB_FREE:管家&省心投放
	} `json:"filtering,omitempty"`                   // 过滤条件，若此字段不传，或传空则视为无限制条件
	OrderField string `json:"order_field,omitempty"` // 排序字段，所有的统计指标均可参与排序
	OrderType  string `json:"order_type,omitempty"`  // 排序方式；默认值: DESC；允许值: ASC, DESC
	Page       int32  `json:"page,omitempty"`        // 页数默认值: 1
	PageSize   int32  `json:"page_size,omitempty"`   // 页面大小默认值:20，page_size范围为[1,1000]
}

func (p ReportCampaignGetParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

// 广告计划数据 - 请求参数
type ReportAdGetParams struct {
	AdvertiserId    string   `json:"advertiser_id,omitempty"`    // 广告主ID
	StartDate       string   `json:"start_date,omitempty"`       // 起始日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期
	EndDate         string   `json:"end_date,omitempty"`         // 结束日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期，时间跨度不能超过30天
	Fields          []string `json:"fields,omitempty"`           // 指定需要的指标名称，可参考应答参数返回的消耗指标字段 默认值：cost、show、avg_show_cost、click、ctr、avg_click_cost、convert、convert_rate、convert_cost
	GroupBy         []string `json:"group_by,omitempty"`         // 分组条件 默认为STAT_GROUP_BY_FIELD_STAT_TIME，支持多种分组条件，具体详见【分组组合规则】
	TimeGranularity string   `json:"time_granularity,omitempty"` // 时间粒度 默认值: STAT_TIME_GRANULARITY_DAILY 允许值:STAT_TIME_GRANULARITY_DAILY (按天维度),STAT_TIME_GRANULARITY_HOURLY (按小时维度)
	Filtering       struct {
		CampaignIds  []int64  `json:"campaign_ids,omitempty"`  // 广告组id列表：按照campaign_id过滤，最多支持100个
		AdName       string   `json:"ad_name,omitempty"`       // 广告计划名称：按照ad_name过滤，字符串长度限制1-30
		AdIds        []int64  `json:"ad_ids,omitempty"`        // 广告计划id列表：按照 ad_id 过滤，最多支持100个
		Pricings     []string `json:"pricings,omitempty"`      // 出价方式列表：按照出价方式过滤，详见【附录-计划出价类型】
		LandingTypes []string `json:"landing_types,omitempty"` // 推广目的列表：按照广告组推广目的过滤,详见【附录-推广目的类型】 允许值:"LINK","APP","DPA","GOODS","STORE","SHOP","AWEME"
		Status       string   `json:"status,omitempty"`        // 广告计划状态：按照广告计划状态过滤，默认为返回“所有不包含已删除”，如果要返回所有包含已删除有对应枚举表示，详见【附录-广告计划投放状态】
		DeliveryMode []string `json:"delivery_mode,omitempty"` // 投放模式筛选。 允许值: STANDARD:标准投放 ADLAB_FREE:管家&省心投放
	} `json:"filtering,omitempty"`                   // 过滤条件，若此字段不传，或传空则视为无限制条件
	OrderField string `json:"order_field,omitempty"` // 排序字段，所有的统计指标均可参与排序
	OrderType  string `json:"order_type,omitempty"`  // 排序方式；默认值: DESC；允许值: ASC, DESC
	Page       int32  `json:"page,omitempty"`        // 页数默认值: 1
	PageSize   int32  `json:"page_size,omitempty"`   // 页面大小默认值:20，page_size范围为[1,1000]
}

func (p ReportAdGetParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

// 广告创意数据 - 请求参数
type ReportCreativeGetParams struct {
	AdvertiserId    string   `json:"advertiser_id,omitempty"`    // 广告主ID
	StartDate       string   `json:"start_date,omitempty"`       // 起始日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期
	EndDate         string   `json:"end_date,omitempty"`         // 结束日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期，时间跨度不能超过30天
	Fields          []string `json:"fields,omitempty"`           // 指定需要的指标名称，可参考应答参数返回的消耗指标字段 默认值：cost、show、avg_show_cost、click、ctr、avg_click_cost、convert、convert_rate、convert_cost
	GroupBy         []string `json:"group_by,omitempty"`         // 分组条件 默认为STAT_GROUP_BY_FIELD_STAT_TIME，支持多种分组条件，具体详见【分组组合规则】
	TimeGranularity string   `json:"time_granularity,omitempty"` // 时间粒度 默认值: STAT_TIME_GRANULARITY_DAILY 允许值:STAT_TIME_GRANULARITY_DAILY (按天维度),STAT_TIME_GRANULARITY_HOURLY (按小时维度)
	Filtering       struct {
		CampaignIds           []int64  `json:"campaign_ids,omitempty"`            // 广告组id列表：按照campaign_id过滤，最多支持100个
		AdIds                 []int64  `json:"ad_ids,omitempty"`                  // 广告计划id列表：按照 ad_id 过滤，最多支持100个
		CreativeIds           []int64  `json:"creative_ids,omitempty"`            // 广告创意id列表：按照 creative_id 过滤，最多支持100个
		InventoryTypes        []string `json:"inventory_types,omitempty"`         // 广告位置列表：按照广告位置过滤，详见【附录-首选投放位置】
		Pricings              []string `json:"pricings,omitempty"`                // 出价方式列表：按照出价方式过滤，详见【附录-计划出价类型】
		ImageModes            []string `json:"image_modes,omitempty"`             // 素材类型列表：按照类型过滤，详见【附录-素材类型】
		CreativeMaterialModes []string `json:"creative_material_modes,omitempty"` // 创意类型列表：按照创意类型过滤，STATIC_ASSEMBLE 表示程序化创意，INTERVENE表示自定义创意。
		LandingTypes          []string `json:"landing_types,omitempty"`           // 推广目的列表：按照广告组推广目的过滤,详见【附录-推广目的类型】 允许值:"LINK","APP","DPA","GOODS","STORE","SHOP","AWEME"
		Status                string   `json:"status,omitempty"`                  // 广告计划状态：按照广告计划状态过滤，默认为返回“所有不包含已删除”，如果要返回所有包含已删除有对应枚举表示，详见【附录-广告计划投放状态】
		DeliveryMode          []string `json:"delivery_mode,omitempty"`           // 投放模式筛选。 允许值: STANDARD:标准投放 ADLAB_FREE:管家&省心投放
	} `json:"filtering,omitempty"`                   // 过滤条件，若此字段不传，或传空则视为无限制条件
	OrderField string `json:"order_field,omitempty"` // 排序字段，所有的统计指标均可参与排序
	OrderType  string `json:"order_type,omitempty"`  // 排序方式；默认值: DESC；允许值: ASC, DESC
	Page       int32  `json:"page,omitempty"`        // 页数默认值: 1
	PageSize   int32  `json:"page_size,omitempty"`   // 页面大小默认值:20，page_size范围为[1,1000]
}

func (p ReportCreativeGetParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}
