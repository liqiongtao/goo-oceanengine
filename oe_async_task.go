package goo_oceanengine

import (
	goo_http_request "github.com/liqiongtao/googo.io/goo-http-request"
	goo_utils "github.com/liqiongtao/googo.io/goo-utils"
)

func AsyncTask(config Config) asyncTask {
	return asyncTask{oceanengine{config: config}}
}

type asyncTask struct {
	oceanengine
}

// 创建异步任务
// https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710562799616#item-link-%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80
// 创建异步任务，每个开发者每天最多只能为每个用户创建 10 个任务（不包括提交失败的任务）。
// 当前支持两种报表类型：普通报表和 DPA 报表
// 注意：
// 1. 暂不支持获取当天的数据
// 2. 暂不支持查询小时纬度分组数据（DPA报表支持）
func (at asyncTask) Create(params AsyncTaskCreateParams, accessToken string) goo_utils.Params {
	return at.get(ASYNC_TASK_CREATE_URL, params.Json(), goo_http_request.HeaderOption("Access-Token", accessToken))
}

// 获取任务列表
func (at asyncTask) Get(params AsyncTaskGetParams, accessToken string) goo_utils.Params {
	return at.get(ASYNC_TASK_GET_URL, params.Json(), goo_http_request.HeaderOption("Access-Token", accessToken))
}

// 下载任务结果
func (at asyncTask) Download(params AsyncTaskDownloadParams, accessToken string) goo_utils.Params {
	return at.get(ASYNC_TASK_DOWNLOAD_URL, params.Json(), goo_http_request.HeaderOption("Access-Token", accessToken))
}
