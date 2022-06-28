package goo_oceanengine

import (
	goo_http_request "github.com/liqiongtao/googo.io/goo-http-request"
)

func DMP(config Config) dmp {
	return dmp{oceanengine{config: config}}
}

type dmp struct {
	oceanengine
}

// 数据源文件上传
// https://open.oceanengine.com/labels/7/docs/1696710568556544
func (d dmp) DataSourceFileUpload(p AdGetParams, accessToken string) (rst DataSourceFileUploadResult) {
	opt := goo_http_request.HeaderOption("Access-Token", accessToken)

	rst = DataSourceFileUploadResult{}
	if err := d.post(DMP_DATASOURCE_FILE_UPLOAD_URL, p.Json(), &rst, opt); err != nil {
		rst = DataSourceFileUploadResult{Code: 5001, Message: err.Error()}
	}
	return
}

// 数据源创建
// https://open.oceanengine.com/labels/7/docs/1696710569089024
func (d dmp) DataSourceCreate(p AdGetParams, accessToken string) (rst DataSourceCreateResult) {
	opt := goo_http_request.HeaderOption("Access-Token", accessToken)

	rst = DataSourceCreateResult{}
	if err := d.get(DMP_DATASOURCE_CREATE_URL, p.Json(), &rst, opt); err != nil {
		rst = DataSourceCreateResult{Code: 5001, Message: err.Error()}
	}
	return
}

// 数据源更新
// https://open.oceanengine.com/labels/7/docs/1696710569591823
func (d dmp) DataSourceUpdate(p AdGetParams, accessToken string) (rst Result) {
	opt := goo_http_request.HeaderOption("Access-Token", accessToken)

	rst = Result{}
	if err := d.get(DMP_DATASOURCE_UPDATE_URL, p.Json(), &rst, opt); err != nil {
		rst = Result{Code: 5001, Message: err.Error()}
	}
	return
}

// 数据源详细信息
// https://open.oceanengine.com/labels/7/docs/1696710570207247
func (d dmp) DataSourceRead(p AdGetParams, accessToken string) (rst DataSourceReadResult) {
	opt := goo_http_request.HeaderOption("Access-Token", accessToken)

	rst = DataSourceReadResult{}
	if err := d.get(DMP_DATASOURCE_READ_URL, p.Json(), &rst, opt); err != nil {
		rst = DataSourceReadResult{Code: 5001, Message: err.Error()}
	}
	return
}

// 人群包列表
// https://open.oceanengine.com/labels/7/docs/1696710570721295
func (d dmp) CustomAudienceSelect(p AdGetParams, accessToken string) (rst CustomAudienceSelectResult) {
	opt := goo_http_request.HeaderOption("Access-Token", accessToken)

	rst = CustomAudienceSelectResult{}
	if err := d.get(DMP_CUSTOM_AUDIENCE_SELECT_URL, p.Json(), &rst, opt); err != nil {
		rst = CustomAudienceSelectResult{Code: 5001, Message: err.Error()}
	}
	return
}

// 人群包详细信息
// https://open.oceanengine.com/labels/7/docs/1696710571259916
func (d dmp) CustomAudienceRead(p AdGetParams, accessToken string) (rst CustomAudienceReadResult) {
	opt := goo_http_request.HeaderOption("Access-Token", accessToken)

	rst = CustomAudienceReadResult{}
	if err := d.get(DMP_CUSTOM_AUDIENCE_READ_URL, p.Json(), &rst, opt); err != nil {
		rst = CustomAudienceReadResult{Code: 5001, Message: err.Error()}
	}
	return
}

// 发布人群包
// https://open.oceanengine.com/labels/7/docs/1696710571768844
func (d dmp) CustomAudiencePublish(p AdGetParams, accessToken string) (rst Result) {
	opt := goo_http_request.HeaderOption("Access-Token", accessToken)

	rst = Result{}
	if err := d.get(DMP_CUSTOM_AUDIENCE_PUBLISH_URL, p.Json(), &rst, opt); err != nil {
		rst = Result{Code: 5001, Message: err.Error()}
	}
	return
}

// 推送人群包
// https://open.oceanengine.com/labels/7/docs/1696710572311552
func (d dmp) CustomAudiencePushV2(p AdGetParams, accessToken string) (rst Result) {
	opt := goo_http_request.HeaderOption("Access-Token", accessToken)

	rst = Result{}
	if err := d.get(DMP_CUSTOM_AUDIENCE_PUSH_V2_URL, p.Json(), &rst, opt); err != nil {
		rst = Result{Code: 5001, Message: err.Error()}
	}
	return
}

// 删除人群包
// https://open.oceanengine.com/labels/7/docs/1696710572836879
func (d dmp) CustomAudienceDelete(p AdGetParams, accessToken string) (rst CustomAudienceDeleteResult) {
	opt := goo_http_request.HeaderOption("Access-Token", accessToken)

	rst = CustomAudienceDeleteResult{}
	if err := d.get(DMP_CUSTOM_AUDIENCE_DELETE_URL, p.Json(), &rst, opt); err != nil {
		rst = CustomAudienceDeleteResult{Code: 5001, Message: err.Error()}
	}
	return
}
