package goo_oceanengine

import "encoding/json"

type Result struct {
	Message string `json:"message" bson:"message"`
	Code    int    `json:"code" bson:"code"`
	Data    struct {
	} `json:"data" bson:"data"`
	RequestId string `json:"request_id" bson:"request_id"`
}

func (p Result) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type DataSourceFileUploadParams struct {
	AdvertiserId int64  `json:"advertiser_id"` // 广告主ID
	File         string `json:"file"`          // 文件
}

type DataSourceFileUploadResult struct {
	Message string `json:"message" bson:"message"`
	Code    int    `json:"code" bson:"code"`
	Data    struct {
		FilePath string `json:"file_path" bson:"file_path"`
	} `json:"data" bson:"data"`
	RequestId string `json:"request_id" bson:"request_id"`
}

func (p DataSourceFileUploadResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type DataSourceCreateParams struct {
	AdvertiserId    int64    `json:"advertiser_id"`     // 广告主ID
	DataSourceName  string   `json:"data_source_name"`  // 数据源名称, 限30个字符内
	Description     string   `json:"description"`       // 数据源描述, 限256个字符内
	DataFormat      int      `json:"data_format"`       // 数据格式, 枚举值:"0"：ProtocolBuffer
	FileStorageType int      `json:"file_storage_type"` // 数据存储类型, 枚举值:"0"：API
	FilePaths       []string `json:"file_paths"`        // 通过【数据源文件上传】接口得到的文件路径，注意：一次上传最多1000个

	// 投放数据源类型，枚举值如下: "UID"：用户ID, "DID"：设备ID ,默认值： "UID"
	// 人群类型仅IMEI、IDFA、IMEI-MD5、IDFA-MD5、OAID、OAID_MD5支持DID投放
	DataSourceType string `json:"data_source_type"`
}

func (p DataSourceCreateParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type DataSourceCreateResult struct {
	Message string `json:"message" bson:"message"`
	Code    int    `json:"code" bson:"code"`
	Data    struct {
		DataSourceId string `json:"data_source_id" bson:"data_source_id"`
	} `json:"data" bson:"data"`
	RequestId string `json:"request_id" bson:"request_id"`
}

func (p DataSourceCreateResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type DataSourceUpdateParams struct {
	AdvertiserId    int64    `json:"advertiser_id"`     // 广告主ID
	DataSourceId    string   `json:"data_source_id"`    // 数据源ID
	OperationType   int      `json:"operation_type"`    // 更新操作类型, 枚举值："1" ：添加,"2"：删除, "3"：重置
	DataFormat      int      `json:"data_format"`       // 数据格式, 枚举值:"0"：ProtocolBuffer
	FileStorageType int      `json:"file_storage_type"` // 数据存储类型, 枚举值:"0"：API
	FilePaths       []string `json:"file_paths"`        // 通过【数据源文件上传】接口得到的文件路径，注意：一次上传最多1000个

	// 投放数据源类型，枚举值如下: "UID"：用户ID, "DID"：设备ID ,默认值： "UID"
	// 人群类型仅IMEI、IDFA、IMEI-MD5、IDFA-MD5、OAID、OAID_MD5支持DID投放
	DataSourceType string `json:"data_source_type"`
}

func (p DataSourceUpdateParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type DataSourceReadParams struct {
	AdvertiserId     int64    `json:"advertiser_id"`       // 广告主ID
	DataSourceIdList []string `json:"data_source_id_list"` // 数据源ID列表 一次最多传400个数据源id
}

func (p DataSourceReadParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type DataSourceReadResult struct {
	Message string `json:"message" bson:"message"`
	Code    int    `json:"code" bson:"code"`
	Data    struct {
		DataList []DataSourceReadListItem `json:"data_list" bson:"data_list"`
	} `json:"data" bson:"data"`
	RequestId string `json:"request_id" bson:"request_id"`
}

func (p DataSourceReadResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type DataSourceReadListItem struct {
	Status      int    `json:"status" bson:"status"`
	Description string `json:"description" bson:"description"`
	ChangeLogs  []struct {
		ChangeLogId   int      `json:"change_log_id" bson:"change_log_id"`
		Status        int      `json:"status" bson:"status"`
		PublishedTime string   `json:"published_time" bson:"published_time"`
		FilePaths     []string `json:"file_paths" bson:"file_paths"`
		CreateTime    string   `json:"create_time" bson:"create_time"`
		ModifyTime    string   `json:"modify_time" bson:"modify_time"`
		Action        int      `json:"action" bson:"action"`
		Message       string   `json:"message" bson:"message"`
	} `json:"change_logs" bson:"change_logs"`
	DefaultAudience struct {
		Status           int    `json:"status" bson:"status"`
		PushStatus       int    `json:"push_status" bson:"push_status"`
		DeliveryStatus   string `json:"delivery_status" bson:"delivery_status"`
		Name             string `json:"name" bson:"name"`
		CustomType       int    `json:"custom_type" bson:"custom_type"`
		UploadNum        int    `json:"upload_num" bson:"upload_num"`
		CustomAudienceId int    `json:"custom_audience_id" bson:"custom_audience_id"`
		CoverNum         int    `json:"cover_num" bson:"cover_num"`
		AdvertiserId     int    `json:"advertiser_id" bson:"advertiser_id"`
		Source           int    `json:"source" bson:"source"`
		ExpiryDate       string `json:"expiry_date" bson:"expiry_date"`
		CreateTime       string `json:"create_time" bson:"create_time"`
		ModifyTime       string `json:"modify_time" bson:"modify_time"`
		Isdel            int    `json:"isdel" bson:"isdel"`
		Tag              string `json:"tag" bson:"tag"`
		CalculateSubType int    `json:"calculate_sub_type" bson:"calculate_sub_type"`
		CalculateType    int    `json:"calculate_type" bson:"calculate_type"`
		DataSourceId     string `json:"data_source_id" bson:"data_source_id"`
		ThirdPartyInfo   string `json:"third_party_info" bson:"third_party_info"`
	} `json:"default_audience" bson:"default_audience"`
	DataSourceType              string `json:"data_source_type" bson:"data_source_type"`
	CreateTime                  string `json:"create_time" bson:"create_time"`
	ModifyTime                  string `json:"modify_time" bson:"modify_time"`
	LastestPublishedChangeLogId int    `json:"lastest_published_change_log_id" bson:"lastest_published_change_log_id"`
	LastestPublishedTime        string `json:"lastest_published_time" bson:"lastest_published_time"`
	UploadNum                   int    `json:"upload_num" bson:"upload_num"`
	CoverNum                    int    `json:"cover_num" bson:"cover_num"`
	DataSourceId                string `json:"data_source_id" bson:"data_source_id"`
	Name                        string `json:"name" bson:"name"`
}

type CustomAudienceSelectParams struct {
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID
	SelectType   int   `json:"select_type"`   // 数据源ID列表 一次最多传400个数据源id
	Offset       int   `json:"offset"`        // 偏移,类似于SQL中offset(起始为0,翻页时new_offset=old_offset+limit），默认值：0，取值范围:≥ 0
	Limit        int   `json:"limit"`         // 返回数据量，默认值：100，取值范围：1-100
}

func (p CustomAudienceSelectParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type CustomAudienceSelectResult struct {
	Message string `json:"message" bson:"message"`
	Code    int    `json:"code" bson:"code"`
	Data    struct {
		CustomAudienceList []CustomAudienceSelectListItem `json:"custom_audience_list" bson:"custom_audience_list"`
		TotalNum           int                            `json:"total_num" bson:"total_num"`
		Offset             int                            `json:"offset" bson:"offset"`
	} `json:"data" bson:"data"`
	RequestId string `json:"request_id" bson:"request_id"`
}

func (p CustomAudienceSelectResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type CustomAudienceSelectListItem struct {
	Status           int    `json:"status" bson:"status"`
	Name             string `json:"name" bson:"name"`
	CustomAudienceId int    `json:"custom_audience_id" bson:"custom_audience_id"`
	Source           string `json:"source" bson:"source"`
	Tag              string `json:"tag" bson:"tag"`
	Isdel            int    `json:"isdel" bson:"isdel"`
	DeliveryStatus   string `json:"delivery_status" bson:"delivery_status"`
	UploadNum        int    `json:"upload_num" bson:"upload_num"`
	CoverNum         int    `json:"cover_num" bson:"cover_num"`
	DataSourceId     string `json:"data_source_id" bson:"data_source_id"`
}

type CustomAudienceReadParams struct {
	AdvertiserId      int64   `json:"advertiser_id"`       // 广告主ID
	CustomAudienceIds []int64 `json:"custom_audience_ids"` // 人群包ID列表，长度取值范围：1-100
}

func (p CustomAudienceReadParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type CustomAudienceReadResult struct {
	Message string `json:"message" bson:"message"`
	Code    int    `json:"code" bson:"code"`
	Data    struct {
		CustomAudienceList []CustomAudienceReadListItem `json:"custom_audience_list" bson:"custom_audience_list"`
	} `json:"data" bson:"data"`
	RequestId string `json:"request_id" bson:"request_id"`
}

func (p CustomAudienceReadResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type CustomAudienceReadListItem struct {
	Status           int    `json:"status" bson:"status"`
	Name             string `json:"name" bson:"name"`
	CustomAudienceId int    `json:"custom_audience_id" bson:"custom_audience_id"`
	CoverNum         int    `json:"cover_num" bson:"cover_num"`
	Source           string `json:"source" bson:"source"`
	Tag              string `json:"tag" bson:"tag"`
	CreateTime       string `json:"create_time" bson:"create_time"`
	ModifyTime       string `json:"modify_time" bson:"modify_time"`
	Isdel            int    `json:"isdel" bson:"isdel"`
	DeliveryStatus   string `json:"delivery_status" bson:"delivery_status"`
	PushStatus       int    `json:"push_status" bson:"push_status"`
	UploadNum        int    `json:"upload_num" bson:"upload_num"`
	ThirdPartyInfo   string `json:"third_party_info" bson:"third_party_info"`
	DataSourceId     string `json:"data_source_id" bson:"data_source_id"`
	ExistPullOffTag  int    `json:"exist_pull_off_tag" bson:"exist_pull_off_tag"`
}

type CustomAudiencePublishParams struct {
	AdvertiserId     int64 `json:"advertiser_id"`      // 广告主ID
	CustomAudienceId int64 `json:"custom_audience_id"` // 人群包ID
}

func (p CustomAudiencePublishParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type CustomAudiencePushV2Params struct {
	AdvertiserId        int64   `json:"advertiser_id"`         // 人群包所属广告主ID
	CustomAudienceId    int64   `json:"custom_audience_id"`    // 人群包ID
	TargetAdvertiserIds []int64 `json:"target_advertiser_ids"` // 推送广告主ID列表，最多推送100个广告主
}

func (p CustomAudiencePushV2Params) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type CustomAudienceDeleteParams struct {
	AdvertiserId     int64 `json:"advertiser_id"`      // 广告主ID
	CustomAudienceId int64 `json:"custom_audience_id"` // 人群包ID
}

func (p CustomAudienceDeleteParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type CustomAudienceDeleteResult struct {
	Message string `json:"message" bson:"message"`
	Code    int    `json:"code" bson:"code"`
	Data    struct {
		CustomAudienceId int `json:"custom_audience_id" bson:"custom_audience_id"`
	} `json:"data" bson:"data"`
	RequestId string `json:"request_id" bson:"request_id"`
}

func (p CustomAudienceDeleteResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}
