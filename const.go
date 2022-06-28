package goo_oceanengine

const (
	// oauth2 - 授权链接
	OAUTH_URL = "https://ad.oceanengine.com/openapi/audit/oauth.html?app_id=%s&state=%s&scope=%s&redirect_uri=%s"
	// oauth2 - 获取access_token
	ACCESS_TOKEN_URL = "https://ad.oceanengine.com/open_api/oauth2/access_token/"
	// oauth2 - 刷新token
	REFRESH_TOKEN_URL = "https://ad.oceanengine.com/open_api/oauth2/refresh_token/"
	// oauth2 - 获取已授权账户
	ADVERTISER_GET_URL = "https://ad.oceanengine.com/open_api/oauth2/advertiser/get/"
	// oauth2 - 获取授权User信息
	USER_INFO_URL = "https://ad.oceanengine.com/open_api/2/user/info/"

	// 账号服务 - 纵横组织账户管理 - 获取纵横组织下资产账户列表
	MAJORDOMO_ADVERTISER_SELECT_URL = "https://ad.oceanengine.com/open_api/2/majordomo/advertiser/select/"

	// 广告投放 - 广告组 - 获取广告组
	CAMPAIGN_GET_URL = "https://ad.oceanengine.com/open_api/2/campaign/get/"

	// 广告投放 - 广告计划模块 - 获取广告计划
	AD_GET_URL = "https://ad.oceanengine.com/open_api/2/ad/get/"

	// 数据报表 - 广告数据报表 - 广告主数据
	REPORT_ADVERTISER_GET_URL = "https://ad.oceanengine.com/open_api/2/report/advertiser/get/"
	// 数据报表 - 广告数据报表 - 广告组数据
	REPORT_CAMPAIGN_GET_URL = "https://ad.oceanengine.com/open_api/2/report/campaign/get/"
	// 数据报表 - 广告数据报表 - 广告计划数据
	REPORT_AD_GET_URL = "https://ad.oceanengine.com/open_api/2/report/ad/get/"
	// 数据报表 - 广告数据报表 - 广告创意数据
	REPORT_CREATIVE_GET_URL = "https://ad.oceanengine.com/open_api/2/report/creative/get/"

	// 数据报表 - 异步数据报表 - 创建异步任务
	ASYNC_TASK_CREATE_URL = "https://ad.oceanengine.com/open_api/2/async_task/create/"
	// 数据报表 - 异步数据报表 - 获取任务列表
	ASYNC_TASK_GET_URL = "https://ad.oceanengine.com/open_api/2/async_task/get/"
	// 数据报表 - 异步数据报表 - 下载任务结果
	ASYNC_TASK_DOWNLOAD_URL = "https://ad.oceanengine.com/open_api/2/async_task/download/"

	// DMP人群管理 - 数据源文件上传
	DMP_DATASOURCE_FILE_UPLOAD_URL = "https://ad.oceanengine.com/open_api/2/dmp/data_source/file/upload/"
	// DMP人群管理 - 数据源创建
	DMP_DATASOURCE_CREATE_URL = "https://ad.oceanengine.com/open_api/2/dmp/data_source/create/"
	// DMP人群管理 - 数据源更新
	DMP_DATASOURCE_UPDATE_URL = "https://ad.oceanengine.com/open_api/2/dmp/data_source/update/"
	// DMP人群管理 - 数据源详细信息
	DMP_DATASOURCE_READ_URL = "https://ad.oceanengine.com/open_api/2/dmp/data_source/read/"
	// DMP人群管理 - 人群包列表
	DMP_CUSTOM_AUDIENCE_SELECT_URL = "https://ad.oceanengine.com/open_api/2/dmp/custom_audience/select/"
	// DMP人群管理 - 人群包详细信息
	DMP_CUSTOM_AUDIENCE_READ_URL = "https://ad.oceanengine.com/open_api/2/dmp/custom_audience/read/"
	// DMP人群管理 - 发布人群包
	DMP_CUSTOM_AUDIENCE_PUBLISH_URL = "https://ad.oceanengine.com/open_api/2/dmp/custom_audience/publish/"
	// DMP人群管理 - 推送人群包
	DMP_CUSTOM_AUDIENCE_PUSH_V2_URL = "https://ad.oceanengine.com/open_api/2/dmp/custom_audience/push_v2/"
	// DMP人群管理 - 删除人群包
	DMP_CUSTOM_AUDIENCE_DELETE_URL = "https://ad.oceanengine.com/open_api/2/dmp/custom_audience/delete/"
)
