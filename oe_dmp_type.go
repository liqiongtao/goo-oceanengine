package goo_oceanengine

type DataSourceFileUploadResult struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		FilePath string `json:"file_path"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type DataSourceCreateResult struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		DataSourceId string `json:"data_source_id"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type DataSourceReadResult struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		DataList []DataSourceListItem `json:"data_list"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type DataSourceListItem struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	ChangeLogs  []struct {
		ChangeLogId   int      `json:"change_log_id"`
		Status        int      `json:"status"`
		PublishedTime string   `json:"published_time"`
		FilePaths     []string `json:"file_paths"`
		CreateTime    string   `json:"create_time"`
		ModifyTime    string   `json:"modify_time"`
		Action        int      `json:"action"`
		Message       string   `json:"message"`
	} `json:"change_logs"`
	DefaultAudience struct {
		Status           int    `json:"status"`
		PushStatus       int    `json:"push_status"`
		DeliveryStatus   string `json:"delivery_status"`
		Name             string `json:"name"`
		CustomType       int    `json:"custom_type"`
		UploadNum        int    `json:"upload_num"`
		CustomAudienceId int    `json:"custom_audience_id"`
		CoverNum         int    `json:"cover_num"`
		AdvertiserId     int    `json:"advertiser_id"`
		Source           int    `json:"source"`
		ExpiryDate       string `json:"expiry_date"`
		CreateTime       string `json:"create_time"`
		ModifyTime       string `json:"modify_time"`
		Isdel            int    `json:"isdel"`
		Tag              string `json:"tag"`
		CalculateSubType int    `json:"calculate_sub_type"`
		CalculateType    int    `json:"calculate_type"`
		DataSourceId     string `json:"data_source_id"`
		ThirdPartyInfo   string `json:"third_party_info"`
	} `json:"default_audience"`
	DataSourceType              string `json:"data_source_type"`
	CreateTime                  string `json:"create_time"`
	ModifyTime                  string `json:"modify_time"`
	LastestPublishedChangeLogId int    `json:"lastest_published_change_log_id"`
	LastestPublishedTime        string `json:"lastest_published_time"`
	UploadNum                   int    `json:"upload_num"`
	CoverNum                    int    `json:"cover_num"`
	DataSourceId                string `json:"data_source_id"`
	Name                        string `json:"name"`
}

type CustomAudienceSelectResult struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		CustomAudienceList []CustomAudienceSelectListItem `json:"custom_audience_list"`
		TotalNum           int                            `json:"total_num"`
		Offset             int                            `json:"offset"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}
type CustomAudienceSelectListItem struct {
	Status           int    `json:"status"`
	Name             string `json:"name"`
	CustomAudienceId int    `json:"custom_audience_id"`
	Source           string `json:"source"`
	Tag              string `json:"tag"`
	Isdel            int    `json:"isdel"`
	DeliveryStatus   string `json:"delivery_status"`
	UploadNum        int    `json:"upload_num"`
	CoverNum         int    `json:"cover_num"`
	DataSourceId     string `json:"data_source_id"`
}

type CustomAudienceReadResult struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		CustomAudienceList []CustomAudienceReadListItem `json:"custom_audience_list"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type CustomAudienceReadListItem struct {
	Status           int    `json:"status"`
	Name             string `json:"name"`
	CustomAudienceId int    `json:"custom_audience_id"`
	CoverNum         int    `json:"cover_num"`
	Source           string `json:"source"`
	Tag              string `json:"tag"`
	CreateTime       string `json:"create_time"`
	ModifyTime       string `json:"modify_time"`
	Isdel            int    `json:"isdel"`
	DeliveryStatus   string `json:"delivery_status"`
	PushStatus       int    `json:"push_status"`
	UploadNum        int    `json:"upload_num"`
	ThirdPartyInfo   string `json:"third_party_info"`
	DataSourceId     string `json:"data_source_id"`
	ExistPullOffTag  int    `json:"exist_pull_off_tag"`
}

type Result struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
	} `json:"data"`
	RequestId string `json:"request_id"`
}

type CustomAudienceDeleteResult struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		CustomAudienceId int `json:"custom_audience_id"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}
