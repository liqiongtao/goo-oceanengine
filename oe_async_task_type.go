package goo_oceanengine

import (
	"encoding/json"
)

// 创建异步任务 - 请求参数
type AsyncTaskCreateParams struct {
	AdvertiserId int64  `json:"advertiser_id,omitempty"` // 广告主id
	TaskName     string `json:"task_name,omitempty"`     // 任务名称。最大长度 25 个字符，不能为空字符。
	Force        string `json:"force,omitempty"`         // true/false。是否强制生成新的任务（不复用之前任务的结果）。
	TaskType     string `json:"task_type,omitempty"`     // 任务类型。可选值：“REPORT”（普通报表），"REPORT_DPA"（DPA 报表），"REPORT_BIDWORD"（关键词/搜索词报表）。
	// 任务的参数。
	TaskParams struct {
		// 起始日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期。
		StartDate string `json:"start_date,omitempty"`

		// 结束日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期。如果分组条件中不包括时间类型，那么时间区间不能超过一年 366 天。
		// 如果分组条件包含STAT_GROUP_BY_TIME_MONTH，那么时间区间不能超过一年 366 天。
		// 如果分组条件包含STAT_GROUP_BY_TIME_WEEK，那么时间区间不能超过三个月 93 天。
		// 如果分组条件包括STAT_GROUP_BY_TIME_DAY，那么时间区间不能超过一个月 31 天。
		EndDate string `json:"end_date,omitempty"`

		// 分组条件，【附录：分组条件】
		// 假设一次查询中共有m个campaign_id，n天数据，有以下三种情况：
		// ①group_by=["STAT_GROUP_BY_TIME_DAY"]表示数据按照天粒度聚合，即本次返回最多为n个数据，表示将m个campaign_id的数据按天各自累计。
		// ②group_by=["STAT_GROUP_BY_CAMPAIGN_ID"]表示按照campaign_id聚合，本次返回最多返回m条数据，即按照m个campaign_id各自累加.
		// ③group_by=["STAT_GROUP_BY_CAMPAIGN_ID", "STAT_GROUP_BY_TIME_DAY"]表示按照时间和id同时聚合，最多返回m * n个数据，返回数据中会同时存在id和时间。
		// 同理如果group_by=["STAT_GROUP_BY_TIME_DAY", "STAT_GROUP_BY_CAMPAIGN_ID","STAT_GROUP_BY_INVENTORY"]表示按照天、广告组id和广告位（或者其他细分维度）同时聚合。
		GroupBy string `json:"group_by,omitempty"`

		OrderField string `json:"order_field,omitempty"` // 排序字段，所有的统计指标均可参与排序

		// 排序方式
		// 默认值: DESC
		// 允许值: "ASC", "DESC"
		OrderType string `json:"order_type,omitempty"`

		// 指定需要的消耗相关指标名称。注意：
		// 1. 如果没有指定，那么只返回支持的默认指标名称；
		// 2. 对于不同的分组条件，支持不同的指标；具体说明可以参考：异步任务报表支持指标
		// 目前还不支持返回如下字段：
		// "attribution_convert_cost",
		// "attribution_deep_convert_cost",
		// "attribution_next_day_open_cnt",
		// "attribution_next_day_open_cost",
		// "attribution_next_day_open_rate"
		Fields []string `json:"fields,omitempty"`

		Filtering struct {
			CampaignId            []int64  `json:"campaign_id,omitempty"`             // 按照campaign_id过滤
			AdId                  []int64  `json:"ad_id,omitempty"`                   // 按照ad_id过滤
			CreativeId            []int64  `json:"creative_id,omitempty"`             // 按照creative_id过滤
			LandingType           []string `json:"landing_type,omitempty"`            // 按照广告组推广目的过滤
			Pricing               []string `json:"pricing,omitempty"`                 // 按照出价方式过滤，详见【附录-计划出价类型】
			InventoryType         []string `json:"inventory_type,omitempty"`          // 按照广告首选位置过滤，详见【附录-首选投放位置】
			CreativeMaterialModes []string `json:"creative_material_modes,omitempty"` // 按照创意类型过滤，STATIC_ASSEMBLE 表示程序化创意，INTERVENE表示自定义创意。
			ConvertType           []string `json:"convert_type,omitempty"`            // 按照转化类型过滤，只支持 dpa 报表。详见【附录-转化类型】
			Platform              []string `json:"platform,omitempty"`                // 按照平台类型过滤，只支持 dpa 报表。详见【附录-受众平台类型】
			Bidword               []string `json:"bidword,omitempty"`                 // 按照搜索词过滤，只支持关键词/搜索词报表，关键词可通过【搜索广告-获取关键词】接口进行获取
			Query                 []string `json:"query,omitempty"`                   // 按照搜索词过滤。搜索词表示用户是通过【搜索词】进行检索到相关广告，可用于分析怎样的搜索词比较常用
			ImageMode             []string `json:"image_mode,omitempty"`              // 按照素材类型过滤，详见【附录-素材类型】
		} `json:"filtering,omitempty"`
	} `json:"task_params,omitempty"`
}

func (p AsyncTaskCreateParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type AsyncTaskCreateResult struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		CreateTime string `json:"create_time"`
		TaskName   string `json:"task_name"`
		TaskId     int    `json:"task_id"`
		TaskStatus string `json:"task_status"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

// 获取任务列表 - 请求参数
type AsyncTaskGetParams struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主id
	Filtering    struct {
		TaskIds  []int64 `json:"task_ids,omitempty"`  // 任务 id。最多 10 个
		TaskName string  `json:"task_name,omitempty"` // 任务名称。
	} `json:"filtering,omitempty"` // 过滤条件，若此字段不传，或传空则视为无限制条件
	Page     int32 `json:"page,omitempty"`      // 页数默认值: 1
	PageSize int32 `json:"page_size,omitempty"` // 页面大小默认值:10，page_size范围为[1,10]
}

func (p AsyncTaskGetParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type AsyncTaskGetResult struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		PageInfo struct {
			TotalNumber int `json:"total_number"`
			Page        int `json:"page"`
			PageSize    int `json:"page_size"`
			TotalPage   int `json:"total_page"`
		} `json:"page_info"`
		List []AsyncTaskInfo `json:"list"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type AsyncTaskInfo struct {
	TaskId     int    `json:"task_id"`
	ErrMsg     string `json:"err_msg"`
	CreateTime string `json:"create_time"`
	TaskName   string `json:"task_name"`
	FileSize   int    `json:"file_size"`
	TaskStatus string `json:"task_status"`
}

// 下载任务结果 - 请求参数
type AsyncTaskDownloadParams struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主id
	TaskId       int64 `json:"task_id,omitempty"`       // 任务 id
	RangeFrom    int32 `json:"range_from,omitempty"`    // 分片下载开始位置（从 0 开始）。闭区间。缺省表示从第一个字节开始下载。
	RangeTo      int32 `json:"range_to,omitempty"`      // 分片下载结束位置。闭区间。缺省表示下载到最后一个字节。特别的， -1 表示到最后一个字节。
}

func (p AsyncTaskDownloadParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}
