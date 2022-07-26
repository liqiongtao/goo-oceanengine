package goo_oceanengine

import (
	"encoding/json"
)

// 广告主数据 - 请求参数
type ReportAdvertiserGetParams struct {
	AdvertiserId    int64    `json:"advertiser_id,omitempty"`    // 广告主ID
	StartDate       string   `json:"start_date,omitempty"`       // 起始日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期
	EndDate         string   `json:"end_date,omitempty"`         // 结束日期,格式YYYY-MM-DD,只支持查询2016-10-26及以后的日期，时间跨度不能超过30天
	Fields          []string `json:"fields,omitempty"`           // 指定需要的指标名称，可参考应答参数返回的消耗指标字段 默认值：cost、show、avg_show_cost、click、ctr、avg_click_cost、convert、convert_rate、convert_cost
	GroupBy         []string `json:"group_by,omitempty"`         // 分组条件 默认为STAT_GROUP_BY_FIELD_STAT_TIME，支持多种分组条件，具体详见【分组组合规则】
	TimeGranularity string   `json:"time_granularity,omitempty"` // 时间粒度 默认值: STAT_TIME_GRANULARITY_DAILY 允许值:STAT_TIME_GRANULARITY_DAILY (按天维度),STAT_TIME_GRANULARITY_HOURLY (按小时维度)
	Filtering       struct {
		DeliveryMode []string `json:"delivery_mode,omitempty"` // 投放模式筛选。 允许值: STANDARD:标准投放 ADLAB_FREE:管家&省心投放
	} `json:"filtering,omitempty"` // 过滤条件，若此字段不传，或传空则视为无限制条件
	Page     int32 `json:"page,omitempty"`      // 页数默认值: 1
	PageSize int32 `json:"page_size,omitempty"` // 页面大小默认值:20，page_size范围为[1,1000]
}

func (p ReportAdvertiserGetParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type ReportAdvertiserGetResult struct {
	Message string `json:"message" bson:"message"`
	Code    int    `json:"code" bson:"code"`
	Data    struct {
		PageInfo struct {
			TotalNumber int `json:"total_number" bson:"total_number"`
			Page        int `json:"page" bson:"page"`
			PageSize    int `json:"page_size" bson:"page_size"`
			TotalPage   int `json:"total_page" bson:"total_page"`
		} `json:"page_info" bson:"page_info"`
		List []ReportAdvertiserInfo `json:"list" bson:"list"`
	} `json:"data" bson:"data"`
	RequestId string `json:"request_id" bson:"request_id"`
}

func (p ReportAdvertiserGetResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type ReportAdvertiserInfo struct {
	Show                              int64   `json:"show" bson:"show"`
	NextDayOpenRate                   float64 `json:"next_day_open_rate" bson:"next_day_open_rate"`
	PreLoanCreditCost                 float64 `json:"pre_loan_credit_cost" bson:"pre_loan_credit_cost"`
	LubanLiveGiftCnt                  int     `json:"luban_live_gift_cnt" bson:"luban_live_gift_cnt"`
	Follow                            int     `json:"follow" bson:"follow"`
	ValidPlayRate                     float64 `json:"valid_play_rate" bson:"valid_play_rate"`
	Redirect                          int     `json:"redirect" bson:"redirect"`
	InstallFinishCost                 float64 `json:"install_finish_cost" bson:"install_finish_cost"`
	ValidPlay                         int     `json:"valid_play" bson:"valid_play"`
	WechatLoginCount                  int     `json:"wechat_login_count" bson:"wechat_login_count"`
	Coupon                            int     `json:"coupon" bson:"coupon"`
	InAppUv                           int     `json:"in_app_uv" bson:"in_app_uv"`
	PayCount                          int     `json:"pay_count" bson:"pay_count"`
	ActiveRegisterCost                float64 `json:"active_register_cost" bson:"active_register_cost"`
	LoanCredit                        int     `json:"loan_credit" bson:"loan_credit"`
	LubanLiveEnterCnt                 int     `json:"luban_live_enter_cnt" bson:"luban_live_enter_cnt"`
	ClickWebsite                      int     `json:"click_website" bson:"click_website"`
	TotalPlay                         int     `json:"total_play" bson:"total_play"`
	WechatPayAmount                   float64 `json:"wechat_pay_amount" bson:"wechat_pay_amount"`
	AttributionWechatFirstPay30DCount int     `json:"attribution_wechat_first_pay_30d_count" bson:"attribution_wechat_first_pay_30_d_count"`
	LoanCreditRate                    float64 `json:"loan_credit_rate" bson:"loan_credit_rate"`
	LubanLiveCommentCnt               int     `json:"luban_live_comment_cnt" bson:"luban_live_comment_cnt"`
	HomeVisited                       int     `json:"home_visited" bson:"home_visited"`
	LubanLivePayOrderCount            int     `json:"luban_live_pay_order_count" bson:"luban_live_pay_order_count"`
	Qq                                int     `json:"qq" bson:"qq"`
	Like                              int     `json:"like" bson:"like"`
	AvgShowCost                       float64 `json:"avg_show_cost" bson:"avg_show_cost"`
	Button                            int     `json:"button" bson:"button"`
	LoanCompletionCost                float64 `json:"loan_completion_cost" bson:"loan_completion_cost"`
	InAppCart                         int     `json:"in_app_cart" bson:"in_app_cart"`
	AttributionGamePay7DCost          float64 `json:"attribution_game_pay_7d_cost" bson:"attribution_game_pay_7_d_cost"`
	AttributionWechatLogin30DCount    int     `json:"attribution_wechat_login_30d_count" bson:"attribution_wechat_login_30_d_count"`
	LubanOrderRoi                     float64 `json:"luban_order_roi" bson:"luban_order_roi"`
	ActiveCost                        float64 `json:"active_cost" bson:"active_cost"`
	AttributionDeepConvertCost        float64 `json:"attribution_deep_convert_cost" bson:"attribution_deep_convert_cost"`
	AttributionDeepConvert            int64   `json:"attribution_deep_convert" bson:"attribution_deep_convert"`
	NextDayOpen                       int     `json:"next_day_open" bson:"next_day_open"`
	InstallFinish                     int     `json:"install_finish" bson:"install_finish"`
	DownloadFinishCost                float64 `json:"download_finish_cost" bson:"download_finish_cost"`
	ConvertCost                       float64 `json:"convert_cost" bson:"convert_cost"`
	ActiveRate                        float64 `json:"active_rate" bson:"active_rate"`
	AdvancedCreativeCouponAddition    int     `json:"advanced_creative_coupon_addition" bson:"advanced_creative_coupon_addition"`
	ActivePayCost                     float64 `json:"active_pay_cost" bson:"active_pay_cost"`
	PoiAddressClick                   int     `json:"poi_address_click" bson:"poi_address_click"`
	Cost                              float64 `json:"cost" bson:"cost"`
	LoanCompletion                    int     `json:"loan_completion" bson:"loan_completion"`
	LubanLiveShareCnt                 int     `json:"luban_live_share_cnt" bson:"luban_live_share_cnt"`
	Download                          int64   `json:"download" bson:"download"`
	WechatLoginCost                   float64 `json:"wechat_login_cost" bson:"wechat_login_cost"`
	Click                             int64   `json:"click" bson:"click"`
	InstallFinishRate                 float64 `json:"install_finish_rate" bson:"install_finish_rate"`
	AttributionWechatFirstPay30DCost  float64 `json:"attribution_wechat_first_pay_30d_cost" bson:"attribution_wechat_first_pay_30_d_cost"`
	Play75FeedBreak                   int     `json:"play_75_feed_break" bson:"play_75_feed_break"`
	Play25FeedBreak                   int     `json:"play_25_feed_break" bson:"play_25_feed_break"`
	DownloadStartCost                 float64 `json:"download_start_cost" bson:"download_start_cost"`
	RedirectToShop                    int     `json:"redirect_to_shop" bson:"redirect_to_shop"`
	WechatFirstPayCount               int     `json:"wechat_first_pay_count" bson:"wechat_first_pay_count"`
	LubanLiveGiftAmount               int     `json:"luban_live_gift_amount" bson:"luban_live_gift_amount"`
	PlayDurationSum                   int     `json:"play_duration_sum" bson:"play_duration_sum"`
	AttributionNextDayOpenCost        float64 `json:"attribution_next_day_open_cost" bson:"attribution_next_day_open_cost"`
	PlayOverRate                      float64 `json:"play_over_rate" bson:"play_over_rate"`
	Play50FeedBreak                   int     `json:"play_50_feed_break" bson:"play_50_feed_break"`
	ClickCallDy                       int     `json:"click_call_dy" bson:"click_call_dy"`
	IesChallengeClick                 int     `json:"ies_challenge_click" bson:"ies_challenge_click"`
	Active                            int64   `json:"active" bson:"active"`
	Lottery                           int     `json:"lottery" bson:"lottery"`
	LoanCompletionRate                float64 `json:"loan_completion_rate" bson:"loan_completion_rate"`
	WifiPlay                          int     `json:"wifi_play" bson:"wifi_play"`
	Convert                           int64   `json:"convert" bson:"convert"`
	Ctr                               float64 `json:"ctr" bson:"ctr"`
	AvgClickCost                      float64 `json:"avg_click_cost" bson:"avg_click_cost"`
	GamePayCost                       int     `json:"game_pay_cost" bson:"game_pay_cost"`
	DeepConvertCost                   float64 `json:"deep_convert_cost" bson:"deep_convert_cost"`
	AttributionWechatPay30DRoi        float64 `json:"attribution_wechat_pay_30d_roi" bson:"attribution_wechat_pay_30_d_roi"`
	Wechat                            int     `json:"wechat" bson:"wechat"`
	View                              int64   `json:"view" bson:"view"`
	AttributionConvertCost            float64 `json:"attribution_convert_cost" bson:"attribution_convert_cost"`
	ConsultEffective                  int     `json:"consult_effective" bson:"consult_effective"`
	Comment                           int     `json:"comment" bson:"comment"`
	LubanLiveClickProductCnt          int     `json:"luban_live_click_product_cnt" bson:"luban_live_click_product_cnt"`
	ConvertRate                       float64 `json:"convert_rate" bson:"convert_rate"`
	ActiveRegisterRate                float64 `json:"active_register_rate" bson:"active_register_rate"`
	ClickLandingPage                  int     `json:"click_landing_page" bson:"click_landing_page"`
	GameAddiction                     int     `json:"game_addiction" bson:"game_addiction"`
	MessageAction                     int     `json:"message_action" bson:"message_action"`
	Message                           int     `json:"message" bson:"message"`
	PoiCollect                        int     `json:"poi_collect" bson:"poi_collect"`
	AttributionNextDayOpenCnt         int     `json:"attribution_next_day_open_cnt" bson:"attribution_next_day_open_cnt"`
	AttributionActivePay7DPerCount    int     `json:"attribution_active_pay_7d_per_count" bson:"attribution_active_pay_7_d_per_count"`
	Vote                              int     `json:"vote" bson:"vote"`
	LubanLivePayOrderStatCost         float64 `json:"luban_live_pay_order_stat_cost" bson:"luban_live_pay_order_stat_cost"`
	DownloadFinishRate                float64 `json:"download_finish_rate" bson:"download_finish_rate"`
	LoanCreditCost                    float64 `json:"loan_credit_cost" bson:"loan_credit_cost"`
	DeepConvertRate                   float64 `json:"deep_convert_rate" bson:"deep_convert_rate"`
	GameAddictionCost                 float64 `json:"game_addiction_cost" bson:"game_addiction_cost"`
	Shopping                          int     `json:"shopping" bson:"shopping"`
	LubanOrderCnt                     int     `json:"luban_order_cnt" bson:"luban_order_cnt"`
	AdvancedCreativePhoneClick        int     `json:"advanced_creative_phone_click" bson:"advanced_creative_phone_click"`
	LubanLiveSlidecartClickCnt        int     `json:"luban_live_slidecart_click_cnt" bson:"luban_live_slidecart_click_cnt"`
	Phone                             int     `json:"phone" bson:"phone"`
	WifiPlayRate                      float64 `json:"wifi_play_rate" bson:"wifi_play_rate"`
	Consult                           int     `json:"consult" bson:"consult"`
	WechatFirstPayCost                float64 `json:"wechat_first_pay_cost" bson:"wechat_first_pay_cost"`
	GameAddictionRate                 float64 `json:"game_addiction_rate" bson:"game_addiction_rate"`
	ActivePayRate                     float64 `json:"active_pay_rate" bson:"active_pay_rate"`
	DownloadStart                     int     `json:"download_start" bson:"download_start"`
	CardShow                          int     `json:"card_show" bson:"card_show"`
	MapSearch                         int     `json:"map_search" bson:"map_search"`
	InAppOrder                        int     `json:"in_app_order" bson:"in_app_order"`
	AttributionConvert                int64   `json:"attribution_convert" bson:"attribution_convert"`
	WechatFirstPayRate                float64 `json:"wechat_first_pay_rate" bson:"wechat_first_pay_rate"`
	AttributionGamePay7DCount         int     `json:"attribution_game_pay_7d_count" bson:"attribution_game_pay_7_d_count"`
	ClickDownload                     int     `json:"click_download" bson:"click_download"`
	PreLoanCredit                     int     `json:"pre_loan_credit" bson:"pre_loan_credit"`
	AttributionWechatLogin30DCost     float64 `json:"attribution_wechat_login_30d_cost" bson:"attribution_wechat_login_30_d_cost"`
	Share                             int     `json:"share" bson:"share"`
	ClickInstall                      int     `json:"click_install" bson:"click_install"`
	AdvancedCreativeFormSubmit        int     `json:"advanced_creative_form_submit" bson:"advanced_creative_form_submit"`
	DeepConvert                       int64   `json:"deep_convert" bson:"deep_convert"`
	LiveFansClubJoinCnt               int     `json:"live_fans_club_join_cnt" bson:"live_fans_club_join_cnt"`
	PhoneConnect                      int     `json:"phone_connect" bson:"phone_connect"`
	Id                                int64   `json:"id" bson:"id"`
	IesMusicClick                     int     `json:"ies_music_click" bson:"ies_music_click"`
	Play100FeedBreak                  int     `json:"play_100_feed_break" bson:"play_100_feed_break"`
	CouponSinglePage                  int     `json:"coupon_single_page" bson:"coupon_single_page"`
	Form                              int     `json:"form" bson:"form"`
	ValidPlayCost                     float64 `json:"valid_play_cost" bson:"valid_play_cost"`
	LiveWatchOneMinuteCount           int     `json:"live_watch_one_minute_count" bson:"live_watch_one_minute_count"`
	LubanOrderStatAmount              float64 `json:"luban_order_stat_amount" bson:"luban_order_stat_amount"`
	DownloadStartRate                 float64 `json:"download_start_rate" bson:"download_start_rate"`
	AttributionNextDayOpenRate        float64 `json:"attribution_next_day_open_rate" bson:"attribution_next_day_open_rate"`
	DownloadFinish                    int     `json:"download_finish" bson:"download_finish"`
	StatDatetime                      string  `json:"stat_datetime" bson:"stat_datetime"`
	AttributionWechatFirstPay30DRate  float64 `json:"attribution_wechat_first_pay_30d_rate" bson:"attribution_wechat_first_pay_30_d_rate"`
	AttributionWechatPay30DAmount     float64 `json:"attribution_wechat_pay_30d_amount" bson:"attribution_wechat_pay_30_d_amount"`
	AdvancedCreativeCounselClick      int     `json:"advanced_creative_counsel_click" bson:"advanced_creative_counsel_click"`
	PhoneEffective                    int     `json:"phone_effective" bson:"phone_effective"`
	InAppDetailUv                     int     `json:"in_app_detail_uv" bson:"in_app_detail_uv"`
	AveragePlayTimePerPlay            float64 `json:"average_play_time_per_play" bson:"average_play_time_per_play"`
	PhoneConfirm                      int     `json:"phone_confirm" bson:"phone_confirm"`
	LocationClick                     int     `json:"location_click" bson:"location_click"`
	GamePayCount                      int     `json:"game_pay_count" bson:"game_pay_count"`
	AdvertiserId                      int64   `json:"advertiser_id" bson:"advertiser_id"`
	ClickShopwindow                   int     `json:"click_shopwindow" bson:"click_shopwindow"`
	Register                          int     `json:"register" bson:"register"`
	InAppPay                          int     `json:"in_app_pay" bson:"in_app_pay"`
	LubanLiveFollowCnt                int     `json:"luban_live_follow_cnt" bson:"luban_live_follow_cnt"`
	NextDayOpenCost                   float64 `json:"next_day_open_cost" bson:"next_day_open_cost"`
	AdvancedCreativeFormClick         int     `json:"advanced_creative_form_click" bson:"advanced_creative_form_click"`
}

func (r ReportAdvertiserGetResult) AdvertiserIdArr() (ids []int64) {
	ids = []int64{}

	idMap := map[int64]struct{}{}
	for _, i := range r.Data.List {
		if _, ok := idMap[i.AdvertiserId]; ok {
			idMap[i.AdvertiserId] = struct{}{}
			ids = append(ids, i.AdvertiserId)
		}
	}

	return
}

// 广告组数据 - 请求参数
type ReportCampaignGetParams struct {
	AdvertiserId    int64    `json:"advertiser_id,omitempty"`    // 广告主ID
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
	} `json:"filtering,omitempty"` // 过滤条件，若此字段不传，或传空则视为无限制条件
	OrderField string `json:"order_field,omitempty"` // 排序字段，所有的统计指标均可参与排序
	OrderType  string `json:"order_type,omitempty"`  // 排序方式；默认值: DESC；允许值: ASC, DESC
	Page       int32  `json:"page,omitempty"`        // 页数默认值: 1
	PageSize   int32  `json:"page_size,omitempty"`   // 页面大小默认值:20，page_size范围为[1,1000]
}

func (p ReportCampaignGetParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type ReportCampaignGetResult struct {
	Message string `json:"message" bson:"message"`
	Code    int    `json:"code" bson:"code"`
	Data    struct {
		PageInfo struct {
			TotalNumber int `json:"total_number" bson:"total_number"`
			Page        int `json:"page" bson:"page"`
			PageSize    int `json:"page_size" bson:"page_size"`
			TotalPage   int `json:"total_page" bson:"total_page"`
		} `json:"page_info" bson:"page_info"`
		List []ReportCampaignInfo `json:"list" bson:"list"`
	} `json:"data" bson:"data"`
	RequestId string `json:"request_id" bson:"request_id"`
}

func (p ReportCampaignGetResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type ReportCampaignInfo struct {
	Show                              int64   `json:"show" bson:"show"`
	NextDayOpenRate                   float64 `json:"next_day_open_rate" bson:"next_day_open_rate"`
	PreLoanCreditCost                 float64 `json:"pre_loan_credit_cost" bson:"pre_loan_credit_cost"`
	LubanLiveGiftCnt                  int     `json:"luban_live_gift_cnt" bson:"luban_live_gift_cnt"`
	Follow                            int     `json:"follow" bson:"follow"`
	ValidPlayRate                     float64 `json:"valid_play_rate" bson:"valid_play_rate"`
	Redirect                          int     `json:"redirect" bson:"redirect"`
	InstallFinishCost                 float64 `json:"install_finish_cost" bson:"install_finish_cost"`
	ValidPlay                         int     `json:"valid_play" bson:"valid_play"`
	WechatLoginCount                  int     `json:"wechat_login_count" bson:"wechat_login_count"`
	Coupon                            int     `json:"coupon" bson:"coupon"`
	InAppUv                           int     `json:"in_app_uv" bson:"in_app_uv"`
	PayCount                          int     `json:"pay_count" bson:"pay_count"`
	ActiveRegisterCost                float64 `json:"active_register_cost" bson:"active_register_cost"`
	LoanCredit                        int     `json:"loan_credit" bson:"loan_credit"`
	LubanLiveEnterCnt                 int     `json:"luban_live_enter_cnt" bson:"luban_live_enter_cnt"`
	ClickWebsite                      int     `json:"click_website" bson:"click_website"`
	TotalPlay                         int     `json:"total_play" bson:"total_play"`
	WechatPayAmount                   float64 `json:"wechat_pay_amount" bson:"wechat_pay_amount"`
	AttributionWechatFirstPay30DCount int     `json:"attribution_wechat_first_pay_30d_count" bson:"attribution_wechat_first_pay_30_d_count"`
	LoanCreditRate                    float64 `json:"loan_credit_rate" bson:"loan_credit_rate"`
	LubanLiveCommentCnt               int     `json:"luban_live_comment_cnt" bson:"luban_live_comment_cnt"`
	HomeVisited                       int     `json:"home_visited" bson:"home_visited"`
	LubanLivePayOrderCount            int     `json:"luban_live_pay_order_count" bson:"luban_live_pay_order_count"`
	Qq                                int     `json:"qq" bson:"qq"`
	Like                              int     `json:"like" bson:"like"`
	AvgShowCost                       float64 `json:"avg_show_cost" bson:"avg_show_cost"`
	Button                            int     `json:"button" bson:"button"`
	LoanCompletionCost                float64 `json:"loan_completion_cost" bson:"loan_completion_cost"`
	InAppCart                         int     `json:"in_app_cart" bson:"in_app_cart"`
	AttributionGamePay7DCost          float64 `json:"attribution_game_pay_7d_cost" bson:"attribution_game_pay_7_d_cost"`
	AttributionWechatLogin30DCount    int     `json:"attribution_wechat_login_30d_count" bson:"attribution_wechat_login_30_d_count"`
	LubanOrderRoi                     float64 `json:"luban_order_roi" bson:"luban_order_roi"`
	ActiveCost                        float64 `json:"active_cost" bson:"active_cost"`
	AttributionDeepConvertCost        float64 `json:"attribution_deep_convert_cost" bson:"attribution_deep_convert_cost"`
	AttributionDeepConvert            int64   `json:"attribution_deep_convert" bson:"attribution_deep_convert"`
	NextDayOpen                       int     `json:"next_day_open" bson:"next_day_open"`
	InstallFinish                     int     `json:"install_finish" bson:"install_finish"`
	DownloadFinishCost                float64 `json:"download_finish_cost" bson:"download_finish_cost"`
	ConvertCost                       float64 `json:"convert_cost" bson:"convert_cost"`
	ActiveRate                        float64 `json:"active_rate" bson:"active_rate"`
	CampaignId                        int64   `json:"campaign_id" bson:"campaign_id"`
	AdvancedCreativeCouponAddition    int     `json:"advanced_creative_coupon_addition" bson:"advanced_creative_coupon_addition"`
	ActivePayCost                     float64 `json:"active_pay_cost" bson:"active_pay_cost"`
	PoiAddressClick                   int     `json:"poi_address_click" bson:"poi_address_click"`
	Cost                              float64 `json:"cost" bson:"cost"`
	LoanCompletion                    int     `json:"loan_completion" bson:"loan_completion"`
	LubanLiveShareCnt                 int     `json:"luban_live_share_cnt" bson:"luban_live_share_cnt"`
	Download                          int64   `json:"download" bson:"download"`
	WechatLoginCost                   float64 `json:"wechat_login_cost" bson:"wechat_login_cost"`
	Click                             int64   `json:"click" bson:"click"`
	InstallFinishRate                 float64 `json:"install_finish_rate" bson:"install_finish_rate"`
	AttributionWechatFirstPay30DCost  float64 `json:"attribution_wechat_first_pay_30d_cost" bson:"attribution_wechat_first_pay_30_d_cost"`
	Play75FeedBreak                   int     `json:"play_75_feed_break" bson:"play_75_feed_break"`
	Play25FeedBreak                   int     `json:"play_25_feed_break" bson:"play_25_feed_break"`
	DownloadStartCost                 float64 `json:"download_start_cost" bson:"download_start_cost"`
	RedirectToShop                    int     `json:"redirect_to_shop" bson:"redirect_to_shop"`
	WechatFirstPayCount               int     `json:"wechat_first_pay_count" bson:"wechat_first_pay_count"`
	LubanLiveGiftAmount               int     `json:"luban_live_gift_amount" bson:"luban_live_gift_amount"`
	PlayDurationSum                   int     `json:"play_duration_sum" bson:"play_duration_sum"`
	AttributionNextDayOpenCost        float64 `json:"attribution_next_day_open_cost" bson:"attribution_next_day_open_cost"`
	PlayOverRate                      float64 `json:"play_over_rate" bson:"play_over_rate"`
	Play50FeedBreak                   int     `json:"play_50_feed_break" bson:"play_50_feed_break"`
	ClickCallDy                       int     `json:"click_call_dy" bson:"click_call_dy"`
	IesChallengeClick                 int     `json:"ies_challenge_click" bson:"ies_challenge_click"`
	Active                            int64   `json:"active" bson:"active"`
	Lottery                           int     `json:"lottery" bson:"lottery"`
	LoanCompletionRate                float64 `json:"loan_completion_rate" bson:"loan_completion_rate"`
	WifiPlay                          int     `json:"wifi_play" bson:"wifi_play"`
	Convert                           int64   `json:"convert" bson:"convert"`
	Ctr                               float64 `json:"ctr" bson:"ctr"`
	AvgClickCost                      float64 `json:"avg_click_cost" bson:"avg_click_cost"`
	GamePayCost                       int     `json:"game_pay_cost" bson:"game_pay_cost"`
	DeepConvertCost                   float64 `json:"deep_convert_cost" bson:"deep_convert_cost"`
	AttributionWechatPay30DRoi        float64 `json:"attribution_wechat_pay_30d_roi" bson:"attribution_wechat_pay_30_d_roi"`
	Wechat                            int     `json:"wechat" bson:"wechat"`
	View                              int64   `json:"view" bson:"view"`
	AttributionConvertCost            float64 `json:"attribution_convert_cost" bson:"attribution_convert_cost"`
	ConsultEffective                  int     `json:"consult_effective" bson:"consult_effective"`
	Comment                           int     `json:"comment" bson:"comment"`
	LubanLiveClickProductCnt          int     `json:"luban_live_click_product_cnt" bson:"luban_live_click_product_cnt"`
	ConvertRate                       float64 `json:"convert_rate" bson:"convert_rate"`
	ActiveRegisterRate                float64 `json:"active_register_rate" bson:"active_register_rate"`
	ClickLandingPage                  int     `json:"click_landing_page" bson:"click_landing_page"`
	GameAddiction                     int     `json:"game_addiction" bson:"game_addiction"`
	MessageAction                     int     `json:"message_action" bson:"message_action"`
	CampaignName                      string  `json:"campaign_name" bson:"campaign_name"`
	Message                           int     `json:"message" bson:"message"`
	PoiCollect                        int     `json:"poi_collect" bson:"poi_collect"`
	AttributionNextDayOpenCnt         int     `json:"attribution_next_day_open_cnt" bson:"attribution_next_day_open_cnt"`
	AttributionActivePay7DPerCount    int     `json:"attribution_active_pay_7d_per_count" bson:"attribution_active_pay_7_d_per_count"`
	Vote                              int     `json:"vote" bson:"vote"`
	LubanLivePayOrderStatCost         float64 `json:"luban_live_pay_order_stat_cost" bson:"luban_live_pay_order_stat_cost"`
	DownloadFinishRate                float64 `json:"download_finish_rate" bson:"download_finish_rate"`
	LoanCreditCost                    float64 `json:"loan_credit_cost" bson:"loan_credit_cost"`
	DeepConvertRate                   float64 `json:"deep_convert_rate" bson:"deep_convert_rate"`
	GameAddictionCost                 float64 `json:"game_addiction_cost" bson:"game_addiction_cost"`
	Shopping                          int     `json:"shopping" bson:"shopping"`
	LubanOrderCnt                     int     `json:"luban_order_cnt" bson:"luban_order_cnt"`
	AdvancedCreativePhoneClick        int     `json:"advanced_creative_phone_click" bson:"advanced_creative_phone_click"`
	LubanLiveSlidecartClickCnt        int     `json:"luban_live_slidecart_click_cnt" bson:"luban_live_slidecart_click_cnt"`
	Phone                             int     `json:"phone" bson:"phone"`
	WifiPlayRate                      float64 `json:"wifi_play_rate" bson:"wifi_play_rate"`
	Consult                           int     `json:"consult" bson:"consult"`
	WechatFirstPayCost                float64 `json:"wechat_first_pay_cost" bson:"wechat_first_pay_cost"`
	GameAddictionRate                 float64 `json:"game_addiction_rate" bson:"game_addiction_rate"`
	ActivePayRate                     float64 `json:"active_pay_rate" bson:"active_pay_rate"`
	DownloadStart                     int     `json:"download_start" bson:"download_start"`
	CardShow                          int     `json:"card_show" bson:"card_show"`
	MapSearch                         int     `json:"map_search" bson:"map_search"`
	InAppOrder                        int     `json:"in_app_order" bson:"in_app_order"`
	AttributionConvert                int64   `json:"attribution_convert" bson:"attribution_convert"`
	WechatFirstPayRate                float64 `json:"wechat_first_pay_rate" bson:"wechat_first_pay_rate"`
	AttributionGamePay7DCount         int     `json:"attribution_game_pay_7d_count" bson:"attribution_game_pay_7_d_count"`
	ClickDownload                     int     `json:"click_download" bson:"click_download"`
	PreLoanCredit                     int     `json:"pre_loan_credit" bson:"pre_loan_credit"`
	AttributionWechatLogin30DCost     float64 `json:"attribution_wechat_login_30d_cost" bson:"attribution_wechat_login_30_d_cost"`
	Share                             int     `json:"share" bson:"share"`
	ClickInstall                      int     `json:"click_install" bson:"click_install"`
	AdvancedCreativeFormSubmit        int     `json:"advanced_creative_form_submit" bson:"advanced_creative_form_submit"`
	DeepConvert                       int64   `json:"deep_convert" bson:"deep_convert"`
	LiveFansClubJoinCnt               int     `json:"live_fans_club_join_cnt" bson:"live_fans_club_join_cnt"`
	PhoneConnect                      int     `json:"phone_connect" bson:"phone_connect"`
	Id                                int64   `json:"id" bson:"id"`
	IesMusicClick                     int     `json:"ies_music_click" bson:"ies_music_click"`
	Play100FeedBreak                  int     `json:"play_100_feed_break" bson:"play_100_feed_break"`
	CouponSinglePage                  int     `json:"coupon_single_page" bson:"coupon_single_page"`
	Form                              int     `json:"form" bson:"form"`
	ValidPlayCost                     float64 `json:"valid_play_cost" bson:"valid_play_cost"`
	LiveWatchOneMinuteCount           int     `json:"live_watch_one_minute_count" bson:"live_watch_one_minute_count"`
	LubanOrderStatAmount              float64 `json:"luban_order_stat_amount" bson:"luban_order_stat_amount"`
	DownloadStartRate                 float64 `json:"download_start_rate" bson:"download_start_rate"`
	AttributionNextDayOpenRate        float64 `json:"attribution_next_day_open_rate" bson:"attribution_next_day_open_rate"`
	DownloadFinish                    int     `json:"download_finish" bson:"download_finish"`
	StatDatetime                      string  `json:"stat_datetime" bson:"stat_datetime"`
	AttributionWechatFirstPay30DRate  float64 `json:"attribution_wechat_first_pay_30d_rate" bson:"attribution_wechat_first_pay_30_d_rate"`
	AttributionWechatPay30DAmount     float64 `json:"attribution_wechat_pay_30d_amount" bson:"attribution_wechat_pay_30_d_amount"`
	AdvancedCreativeCounselClick      int     `json:"advanced_creative_counsel_click" bson:"advanced_creative_counsel_click"`
	PhoneEffective                    int     `json:"phone_effective" bson:"phone_effective"`
	InAppDetailUv                     int     `json:"in_app_detail_uv" bson:"in_app_detail_uv"`
	AveragePlayTimePerPlay            float64 `json:"average_play_time_per_play" bson:"average_play_time_per_play"`
	PhoneConfirm                      int     `json:"phone_confirm" bson:"phone_confirm"`
	LocationClick                     int     `json:"location_click" bson:"location_click"`
	GamePayCount                      int     `json:"game_pay_count" bson:"game_pay_count"`
	ClickShopwindow                   int     `json:"click_shopwindow" bson:"click_shopwindow"`
	Register                          int     `json:"register" bson:"register"`
	InAppPay                          int     `json:"in_app_pay" bson:"in_app_pay"`
	LubanLiveFollowCnt                int     `json:"luban_live_follow_cnt" bson:"luban_live_follow_cnt"`
	NextDayOpenCost                   float64 `json:"next_day_open_cost" bson:"next_day_open_cost"`
	AdvancedCreativeFormClick         int     `json:"advanced_creative_form_click" bson:"advanced_creative_form_click"`
}

func (r ReportCampaignGetResult) CampaignIdArr() (ids []int64) {
	ids = []int64{}

	idMap := map[int64]struct{}{}
	for _, i := range r.Data.List {
		if _, ok := idMap[i.CampaignId]; ok {
			idMap[i.CampaignId] = struct{}{}
			ids = append(ids, i.CampaignId)
		}
	}

	return
}

// 广告计划数据 - 请求参数
type ReportAdGetParams struct {
	AdvertiserId    int64    `json:"advertiser_id,omitempty"`    // 广告主ID
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
	} `json:"filtering,omitempty"` // 过滤条件，若此字段不传，或传空则视为无限制条件
	OrderField string `json:"order_field,omitempty"` // 排序字段，所有的统计指标均可参与排序
	OrderType  string `json:"order_type,omitempty"`  // 排序方式；默认值: DESC；允许值: ASC, DESC
	Page       int32  `json:"page,omitempty"`        // 页数默认值: 1
	PageSize   int32  `json:"page_size,omitempty"`   // 页面大小默认值:20，page_size范围为[1,1000]
}

func (p ReportAdGetParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type ReportAdGetResult struct {
	Message string `json:"message" bson:"message"`
	Code    int    `json:"code" bson:"code"`
	Data    struct {
		PageInfo struct {
			TotalNumber int `json:"total_number" bson:"total_number"`
			Page        int `json:"page" bson:"page"`
			PageSize    int `json:"page_size" bson:"page_size"`
			TotalPage   int `json:"total_page" bson:"total_page"`
		} `json:"page_info" bson:"page_info"`
		List []ReportAdInfo `json:"list" bson:"list"`
	} `json:"data" bson:"data"`
	RequestId string `json:"request_id" bson:"request_id"`
}

type ReportAdInfo struct {
	Show                              int64   `json:"show" bson:"show"`
	NextDayOpenRate                   float64 `json:"next_day_open_rate" bson:"next_day_open_rate"`
	PreLoanCreditCost                 float64 `json:"pre_loan_credit_cost" bson:"pre_loan_credit_cost"`
	LubanLiveGiftCnt                  int     `json:"luban_live_gift_cnt" bson:"luban_live_gift_cnt"`
	Follow                            int     `json:"follow" bson:"follow"`
	ValidPlayRate                     float64 `json:"valid_play_rate" bson:"valid_play_rate"`
	Redirect                          int     `json:"redirect" bson:"redirect"`
	InstallFinishCost                 float64 `json:"install_finish_cost" bson:"install_finish_cost"`
	ValidPlay                         int     `json:"valid_play" bson:"valid_play"`
	WechatLoginCount                  int     `json:"wechat_login_count" bson:"wechat_login_count"`
	Coupon                            int     `json:"coupon" bson:"coupon"`
	InAppUv                           int     `json:"in_app_uv" bson:"in_app_uv"`
	PayCount                          int     `json:"pay_count" bson:"pay_count"`
	ActiveRegisterCost                float64 `json:"active_register_cost" bson:"active_register_cost"`
	LoanCredit                        int     `json:"loan_credit" bson:"loan_credit"`
	LubanLiveEnterCnt                 int     `json:"luban_live_enter_cnt" bson:"luban_live_enter_cnt"`
	ClickWebsite                      int     `json:"click_website" bson:"click_website"`
	TotalPlay                         int     `json:"total_play" bson:"total_play"`
	WechatPayAmount                   float64 `json:"wechat_pay_amount" bson:"wechat_pay_amount"`
	AttributionWechatFirstPay30DCount int     `json:"attribution_wechat_first_pay_30d_count" bson:"attribution_wechat_first_pay_30_d_count"`
	LoanCreditRate                    float64 `json:"loan_credit_rate" bson:"loan_credit_rate"`
	AdName                            string  `json:"ad_name" bson:"ad_name"`
	LubanLiveCommentCnt               int     `json:"luban_live_comment_cnt" bson:"luban_live_comment_cnt"`
	HomeVisited                       int     `json:"home_visited" bson:"home_visited"`
	LubanLivePayOrderCount            int     `json:"luban_live_pay_order_count" bson:"luban_live_pay_order_count"`
	Qq                                int     `json:"qq" bson:"qq"`
	Like                              int     `json:"like" bson:"like"`
	AvgShowCost                       float64 `json:"avg_show_cost" bson:"avg_show_cost"`
	Button                            int     `json:"button" bson:"button"`
	LoanCompletionCost                float64 `json:"loan_completion_cost" bson:"loan_completion_cost"`
	InAppCart                         int     `json:"in_app_cart" bson:"in_app_cart"`
	AttributionGamePay7DCost          float64 `json:"attribution_game_pay_7d_cost" bson:"attribution_game_pay_7_d_cost"`
	AttributionWechatLogin30DCount    int     `json:"attribution_wechat_login_30d_count" bson:"attribution_wechat_login_30_d_count"`
	LubanOrderRoi                     float64 `json:"luban_order_roi" bson:"luban_order_roi"`
	ActiveCost                        float64 `json:"active_cost" bson:"active_cost"`
	AttributionDeepConvertCost        float64 `json:"attribution_deep_convert_cost" bson:"attribution_deep_convert_cost"`
	AttributionDeepConvert            int64   `json:"attribution_deep_convert" bson:"attribution_deep_convert"`
	NextDayOpen                       int     `json:"next_day_open" bson:"next_day_open"`
	InstallFinish                     int     `json:"install_finish" bson:"install_finish"`
	DownloadFinishCost                float64 `json:"download_finish_cost" bson:"download_finish_cost"`
	ConvertCost                       float64 `json:"convert_cost" bson:"convert_cost"`
	ActiveRate                        float64 `json:"active_rate" bson:"active_rate"`
	CampaignId                        int64   `json:"campaign_id" bson:"campaign_id"`
	AdvancedCreativeCouponAddition    int     `json:"advanced_creative_coupon_addition" bson:"advanced_creative_coupon_addition"`
	ActivePayCost                     float64 `json:"active_pay_cost" bson:"active_pay_cost"`
	PoiAddressClick                   int     `json:"poi_address_click" bson:"poi_address_click"`
	Cost                              float64 `json:"cost" bson:"cost"`
	LoanCompletion                    int     `json:"loan_completion" bson:"loan_completion"`
	LubanLiveShareCnt                 int     `json:"luban_live_share_cnt" bson:"luban_live_share_cnt"`
	Download                          int64   `json:"download" bson:"download"`
	WechatLoginCost                   float64 `json:"wechat_login_cost" bson:"wechat_login_cost"`
	Click                             int64   `json:"click" bson:"click"`
	InstallFinishRate                 float64 `json:"install_finish_rate" bson:"install_finish_rate"`
	AttributionWechatFirstPay30DCost  float64 `json:"attribution_wechat_first_pay_30d_cost" bson:"attribution_wechat_first_pay_30_d_cost"`
	Play75FeedBreak                   int     `json:"play_75_feed_break" bson:"play_75_feed_break"`
	Play25FeedBreak                   int     `json:"play_25_feed_break" bson:"play_25_feed_break"`
	DownloadStartCost                 float64 `json:"download_start_cost" bson:"download_start_cost"`
	RedirectToShop                    int     `json:"redirect_to_shop" bson:"redirect_to_shop"`
	WechatFirstPayCount               int     `json:"wechat_first_pay_count" bson:"wechat_first_pay_count"`
	LubanLiveGiftAmount               int     `json:"luban_live_gift_amount" bson:"luban_live_gift_amount"`
	PlayDurationSum                   int     `json:"play_duration_sum" bson:"play_duration_sum"`
	AttributionNextDayOpenCost        float64 `json:"attribution_next_day_open_cost" bson:"attribution_next_day_open_cost"`
	PlayOverRate                      float64 `json:"play_over_rate" bson:"play_over_rate"`
	Play50FeedBreak                   int     `json:"play_50_feed_break" bson:"play_50_feed_break"`
	AdId                              int64   `json:"ad_id" bson:"ad_id"`
	ClickCallDy                       int     `json:"click_call_dy" bson:"click_call_dy"`
	IesChallengeClick                 int     `json:"ies_challenge_click" bson:"ies_challenge_click"`
	Active                            int64   `json:"active" bson:"active"`
	Lottery                           int     `json:"lottery" bson:"lottery"`
	LoanCompletionRate                float64 `json:"loan_completion_rate" bson:"loan_completion_rate"`
	WifiPlay                          int     `json:"wifi_play" bson:"wifi_play"`
	Convert                           int64   `json:"convert" bson:"convert"`
	Ctr                               float64 `json:"ctr" bson:"ctr"`
	AvgClickCost                      float64 `json:"avg_click_cost" bson:"avg_click_cost"`
	GamePayCost                       int     `json:"game_pay_cost" bson:"game_pay_cost"`
	DeepConvertCost                   float64 `json:"deep_convert_cost" bson:"deep_convert_cost"`
	AttributionWechatPay30DRoi        float64 `json:"attribution_wechat_pay_30d_roi" bson:"attribution_wechat_pay_30_d_roi"`
	Wechat                            int     `json:"wechat" bson:"wechat"`
	View                              int64   `json:"view" bson:"view"`
	AttributionConvertCost            float64 `json:"attribution_convert_cost" bson:"attribution_convert_cost"`
	ConsultEffective                  int     `json:"consult_effective" bson:"consult_effective"`
	Comment                           int     `json:"comment" bson:"comment"`
	LubanLiveClickProductCnt          int     `json:"luban_live_click_product_cnt" bson:"luban_live_click_product_cnt"`
	ConvertRate                       float64 `json:"convert_rate" bson:"convert_rate"`
	ActiveRegisterRate                float64 `json:"active_register_rate" bson:"active_register_rate"`
	ClickLandingPage                  int     `json:"click_landing_page" bson:"click_landing_page"`
	GameAddiction                     int     `json:"game_addiction" bson:"game_addiction"`
	MessageAction                     int     `json:"message_action" bson:"message_action"`
	CampaignName                      string  `json:"campaign_name" bson:"campaign_name"`
	Message                           int     `json:"message" bson:"message"`
	PoiCollect                        int     `json:"poi_collect" bson:"poi_collect"`
	AttributionNextDayOpenCnt         int     `json:"attribution_next_day_open_cnt" bson:"attribution_next_day_open_cnt"`
	AttributionActivePay7DPerCount    int     `json:"attribution_active_pay_7d_per_count" bson:"attribution_active_pay_7_d_per_count"`
	Vote                              int     `json:"vote" bson:"vote"`
	LubanLivePayOrderStatCost         float64 `json:"luban_live_pay_order_stat_cost" bson:"luban_live_pay_order_stat_cost"`
	DownloadFinishRate                float64 `json:"download_finish_rate" bson:"download_finish_rate"`
	LoanCreditCost                    float64 `json:"loan_credit_cost" bson:"loan_credit_cost"`
	DeepConvertRate                   float64 `json:"deep_convert_rate" bson:"deep_convert_rate"`
	GameAddictionCost                 float64 `json:"game_addiction_cost" bson:"game_addiction_cost"`
	Shopping                          int     `json:"shopping" bson:"shopping"`
	LubanOrderCnt                     int     `json:"luban_order_cnt" bson:"luban_order_cnt"`
	AdvancedCreativePhoneClick        int     `json:"advanced_creative_phone_click" bson:"advanced_creative_phone_click"`
	LubanLiveSlidecartClickCnt        int     `json:"luban_live_slidecart_click_cnt" bson:"luban_live_slidecart_click_cnt"`
	Phone                             int     `json:"phone" bson:"phone"`
	WifiPlayRate                      float64 `json:"wifi_play_rate" bson:"wifi_play_rate"`
	Consult                           int     `json:"consult" bson:"consult"`
	WechatFirstPayCost                float64 `json:"wechat_first_pay_cost" bson:"wechat_first_pay_cost"`
	GameAddictionRate                 float64 `json:"game_addiction_rate" bson:"game_addiction_rate"`
	ActivePayRate                     float64 `json:"active_pay_rate" bson:"active_pay_rate"`
	DownloadStart                     int     `json:"download_start" bson:"download_start"`
	CardShow                          int     `json:"card_show" bson:"card_show"`
	MapSearch                         int     `json:"map_search" bson:"map_search"`
	InAppOrder                        int     `json:"in_app_order" bson:"in_app_order"`
	AttributionConvert                int64   `json:"attribution_convert" bson:"attribution_convert"`
	WechatFirstPayRate                float64 `json:"wechat_first_pay_rate" bson:"wechat_first_pay_rate"`
	AttributionGamePay7DCount         int     `json:"attribution_game_pay_7d_count" bson:"attribution_game_pay_7_d_count"`
	ClickDownload                     int     `json:"click_download" bson:"click_download"`
	PreLoanCredit                     int     `json:"pre_loan_credit" bson:"pre_loan_credit"`
	AttributionWechatLogin30DCost     float64 `json:"attribution_wechat_login_30d_cost" bson:"attribution_wechat_login_30_d_cost"`
	Share                             int     `json:"share" bson:"share"`
	ClickInstall                      int     `json:"click_install" bson:"click_install"`
	AdvancedCreativeFormSubmit        int     `json:"advanced_creative_form_submit" bson:"advanced_creative_form_submit"`
	DeepConvert                       int64   `json:"deep_convert" bson:"deep_convert"`
	LiveFansClubJoinCnt               int     `json:"live_fans_club_join_cnt" bson:"live_fans_club_join_cnt"`
	PhoneConnect                      int     `json:"phone_connect" bson:"phone_connect"`
	Id                                int64   `json:"id" bson:"id"`
	IesMusicClick                     int     `json:"ies_music_click" bson:"ies_music_click"`
	Play100FeedBreak                  int     `json:"play_100_feed_break" bson:"play_100_feed_break"`
	CouponSinglePage                  int     `json:"coupon_single_page" bson:"coupon_single_page"`
	Form                              int     `json:"form" bson:"form"`
	ValidPlayCost                     float64 `json:"valid_play_cost" bson:"valid_play_cost"`
	LiveWatchOneMinuteCount           int     `json:"live_watch_one_minute_count" bson:"live_watch_one_minute_count"`
	LubanOrderStatAmount              float64 `json:"luban_order_stat_amount" bson:"luban_order_stat_amount"`
	DownloadStartRate                 float64 `json:"download_start_rate" bson:"download_start_rate"`
	AttributionNextDayOpenRate        float64 `json:"attribution_next_day_open_rate" bson:"attribution_next_day_open_rate"`
	DownloadFinish                    int     `json:"download_finish" bson:"download_finish"`
	StatDatetime                      string  `json:"stat_datetime" bson:"stat_datetime"`
	AttributionWechatFirstPay30DRate  float64 `json:"attribution_wechat_first_pay_30d_rate" bson:"attribution_wechat_first_pay_30_d_rate"`
	AttributionWechatPay30DAmount     float64 `json:"attribution_wechat_pay_30d_amount" bson:"attribution_wechat_pay_30_d_amount"`
	AdvancedCreativeCounselClick      int     `json:"advanced_creative_counsel_click" bson:"advanced_creative_counsel_click"`
	PhoneEffective                    int     `json:"phone_effective" bson:"phone_effective"`
	InAppDetailUv                     int     `json:"in_app_detail_uv" bson:"in_app_detail_uv"`
	AveragePlayTimePerPlay            float64 `json:"average_play_time_per_play" bson:"average_play_time_per_play"`
	PhoneConfirm                      int     `json:"phone_confirm" bson:"phone_confirm"`
	LocationClick                     int     `json:"location_click" bson:"location_click"`
	GamePayCount                      int     `json:"game_pay_count" bson:"game_pay_count"`
	ClickShopwindow                   int     `json:"click_shopwindow" bson:"click_shopwindow"`
	Register                          int     `json:"register" bson:"register"`
	InAppPay                          int     `json:"in_app_pay" bson:"in_app_pay"`
	LubanLiveFollowCnt                int     `json:"luban_live_follow_cnt" bson:"luban_live_follow_cnt"`
	NextDayOpenCost                   float64 `json:"next_day_open_cost" bson:"next_day_open_cost"`
	AdvancedCreativeFormClick         int     `json:"advanced_creative_form_click" bson:"advanced_creative_form_click"`
}

func (r ReportAdGetResult) CampaignIdArr() (ids []int64) {
	ids = []int64{}

	idMap := map[int64]struct{}{}
	for _, i := range r.Data.List {
		if _, ok := idMap[i.CampaignId]; ok {
			idMap[i.CampaignId] = struct{}{}
			ids = append(ids, i.CampaignId)
		}
	}

	return
}

func (r ReportAdGetResult) AdIdArr() (ids []int64) {
	ids = []int64{}

	idMap := map[int64]struct{}{}
	for _, i := range r.Data.List {
		if _, ok := idMap[i.AdId]; ok {
			idMap[i.AdId] = struct{}{}
			ids = append(ids, i.AdId)
		}
	}

	return
}

// 广告创意数据 - 请求参数
type ReportCreativeGetParams struct {
	AdvertiserId    int64    `json:"advertiser_id,omitempty"`    // 广告主ID
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
	} `json:"filtering,omitempty"` // 过滤条件，若此字段不传，或传空则视为无限制条件
	OrderField string `json:"order_field,omitempty"` // 排序字段，所有的统计指标均可参与排序
	OrderType  string `json:"order_type,omitempty"`  // 排序方式；默认值: DESC；允许值: ASC, DESC
	Page       int32  `json:"page,omitempty"`        // 页数默认值: 1
	PageSize   int32  `json:"page_size,omitempty"`   // 页面大小默认值:20，page_size范围为[1,1000]
}

func (p ReportCreativeGetParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type ReportCreativeGetResult struct {
	Message string `json:"message" bson:"message"`
	Code    int    `json:"code" bson:"code"`
	Data    struct {
		PageInfo struct {
			TotalNumber int `json:"total_number" bson:"total_number"`
			Page        int `json:"page" bson:"page"`
			PageSize    int `json:"page_size" bson:"page_size"`
			TotalPage   int `json:"total_page" bson:"total_page"`
		} `json:"page_info" bson:"page_info"`
		List []ReportCreativeInfo `json:"list" bson:"list"`
	} `json:"data" bson:"data"`
	RequestId string `json:"request_id" bson:"request_id"`
}

type ReportCreativeInfo struct {
	Show                              int64   `json:"show" bson:"show"`
	NextDayOpenRate                   float64 `json:"next_day_open_rate" bson:"next_day_open_rate"`
	PreLoanCreditCost                 float64 `json:"pre_loan_credit_cost" bson:"pre_loan_credit_cost"`
	LubanLiveGiftCnt                  int     `json:"luban_live_gift_cnt" bson:"luban_live_gift_cnt"`
	Follow                            int     `json:"follow" bson:"follow"`
	ValidPlayRate                     float64 `json:"valid_play_rate" bson:"valid_play_rate"`
	Redirect                          int     `json:"redirect" bson:"redirect"`
	InstallFinishCost                 float64 `json:"install_finish_cost" bson:"install_finish_cost"`
	ValidPlay                         int     `json:"valid_play" bson:"valid_play"`
	WechatLoginCount                  int     `json:"wechat_login_count" bson:"wechat_login_count"`
	Coupon                            int     `json:"coupon" bson:"coupon"`
	InAppUv                           int     `json:"in_app_uv" bson:"in_app_uv"`
	PayCount                          int     `json:"pay_count" bson:"pay_count"`
	ActiveRegisterCost                float64 `json:"active_register_cost" bson:"active_register_cost"`
	LoanCredit                        int     `json:"loan_credit" bson:"loan_credit"`
	LubanLiveEnterCnt                 int     `json:"luban_live_enter_cnt" bson:"luban_live_enter_cnt"`
	ClickWebsite                      int     `json:"click_website" bson:"click_website"`
	TotalPlay                         int     `json:"total_play" bson:"total_play"`
	WechatPayAmount                   float64 `json:"wechat_pay_amount" bson:"wechat_pay_amount"`
	AttributionWechatFirstPay30DCount int     `json:"attribution_wechat_first_pay_30d_count" bson:"attribution_wechat_first_pay_30_d_count"`
	LoanCreditRate                    float64 `json:"loan_credit_rate" bson:"loan_credit_rate"`
	AdName                            string  `json:"ad_name" bson:"ad_name"`
	LubanLiveCommentCnt               int     `json:"luban_live_comment_cnt" bson:"luban_live_comment_cnt"`
	HomeVisited                       int     `json:"home_visited" bson:"home_visited"`
	LubanLivePayOrderCount            int     `json:"luban_live_pay_order_count" bson:"luban_live_pay_order_count"`
	Qq                                int     `json:"qq" bson:"qq"`
	Like                              int     `json:"like" bson:"like"`
	AvgShowCost                       float64 `json:"avg_show_cost" bson:"avg_show_cost"`
	Button                            int     `json:"button" bson:"button"`
	LoanCompletionCost                float64 `json:"loan_completion_cost" bson:"loan_completion_cost"`
	InAppCart                         int     `json:"in_app_cart" bson:"in_app_cart"`
	AttributionGamePay7DCost          float64 `json:"attribution_game_pay_7d_cost" bson:"attribution_game_pay_7_d_cost"`
	AttributionWechatLogin30DCount    int     `json:"attribution_wechat_login_30d_count" bson:"attribution_wechat_login_30_d_count"`
	LubanOrderRoi                     float64 `json:"luban_order_roi" bson:"luban_order_roi"`
	ActiveCost                        float64 `json:"active_cost" bson:"active_cost"`
	AttributionDeepConvertCost        float64 `json:"attribution_deep_convert_cost" bson:"attribution_deep_convert_cost"`
	AttributionDeepConvert            int64   `json:"attribution_deep_convert" bson:"attribution_deep_convert"`
	NextDayOpen                       int     `json:"next_day_open" bson:"next_day_open"`
	InstallFinish                     int     `json:"install_finish" bson:"install_finish"`
	DownloadFinishCost                float64 `json:"download_finish_cost" bson:"download_finish_cost"`
	ConvertCost                       float64 `json:"convert_cost" bson:"convert_cost"`
	ActiveRate                        float64 `json:"active_rate" bson:"active_rate"`
	CampaignId                        int64   `json:"campaign_id" bson:"campaign_id"`
	AdvancedCreativeCouponAddition    int     `json:"advanced_creative_coupon_addition" bson:"advanced_creative_coupon_addition"`
	ActivePayCost                     float64 `json:"active_pay_cost" bson:"active_pay_cost"`
	PoiAddressClick                   int     `json:"poi_address_click" bson:"poi_address_click"`
	Cost                              float64 `json:"cost" bson:"cost"`
	LoanCompletion                    int     `json:"loan_completion" bson:"loan_completion"`
	LubanLiveShareCnt                 int     `json:"luban_live_share_cnt" bson:"luban_live_share_cnt"`
	Download                          int64   `json:"download" bson:"download"`
	WechatLoginCost                   float64 `json:"wechat_login_cost" bson:"wechat_login_cost"`
	Click                             int64   `json:"click" bson:"click"`
	InstallFinishRate                 float64 `json:"install_finish_rate" bson:"install_finish_rate"`
	AttributionWechatFirstPay30DCost  float64 `json:"attribution_wechat_first_pay_30d_cost" bson:"attribution_wechat_first_pay_30_d_cost"`
	Play75FeedBreak                   int     `json:"play_75_feed_break" bson:"play_75_feed_break"`
	Play25FeedBreak                   int     `json:"play_25_feed_break" bson:"play_25_feed_break"`
	CreativeId                        int64   `json:"creative_id" bson:"creative_id"`
	DownloadStartCost                 float64 `json:"download_start_cost" bson:"download_start_cost"`
	RedirectToShop                    int     `json:"redirect_to_shop" bson:"redirect_to_shop"`
	WechatFirstPayCount               int     `json:"wechat_first_pay_count" bson:"wechat_first_pay_count"`
	LubanLiveGiftAmount               int     `json:"luban_live_gift_amount" bson:"luban_live_gift_amount"`
	PlayDurationSum                   int     `json:"play_duration_sum" bson:"play_duration_sum"`
	AttributionNextDayOpenCost        float64 `json:"attribution_next_day_open_cost" bson:"attribution_next_day_open_cost"`
	PlayOverRate                      float64 `json:"play_over_rate" bson:"play_over_rate"`
	Play50FeedBreak                   int     `json:"play_50_feed_break" bson:"play_50_feed_break"`
	AdId                              int64   `json:"ad_id" bson:"ad_id"`
	ClickCallDy                       int     `json:"click_call_dy" bson:"click_call_dy"`
	IesChallengeClick                 int     `json:"ies_challenge_click" bson:"ies_challenge_click"`
	Active                            int64   `json:"active" bson:"active"`
	Lottery                           int     `json:"lottery" bson:"lottery"`
	LoanCompletionRate                float64 `json:"loan_completion_rate" bson:"loan_completion_rate"`
	WifiPlay                          int     `json:"wifi_play" bson:"wifi_play"`
	Convert                           int64   `json:"convert" bson:"convert"`
	Ctr                               float64 `json:"ctr" bson:"ctr"`
	AvgClickCost                      float64 `json:"avg_click_cost" bson:"avg_click_cost"`
	GamePayCost                       int     `json:"game_pay_cost" bson:"game_pay_cost"`
	DeepConvertCost                   float64 `json:"deep_convert_cost" bson:"deep_convert_cost"`
	AttributionWechatPay30DRoi        float64 `json:"attribution_wechat_pay_30d_roi" bson:"attribution_wechat_pay_30_d_roi"`
	Wechat                            int     `json:"wechat" bson:"wechat"`
	View                              int64   `json:"view" bson:"view"`
	AttributionConvertCost            float64 `json:"attribution_convert_cost" bson:"attribution_convert_cost"`
	ConsultEffective                  int     `json:"consult_effective" bson:"consult_effective"`
	Comment                           int     `json:"comment" bson:"comment"`
	LubanLiveClickProductCnt          int     `json:"luban_live_click_product_cnt" bson:"luban_live_click_product_cnt"`
	ConvertRate                       float64 `json:"convert_rate" bson:"convert_rate"`
	ActiveRegisterRate                float64 `json:"active_register_rate" bson:"active_register_rate"`
	ClickLandingPage                  int     `json:"click_landing_page" bson:"click_landing_page"`
	GameAddiction                     int     `json:"game_addiction" bson:"game_addiction"`
	MessageAction                     int     `json:"message_action" bson:"message_action"`
	CampaignName                      string  `json:"campaign_name" bson:"campaign_name"`
	Message                           int     `json:"message" bson:"message"`
	PoiCollect                        int     `json:"poi_collect" bson:"poi_collect"`
	AttributionNextDayOpenCnt         int     `json:"attribution_next_day_open_cnt" bson:"attribution_next_day_open_cnt"`
	AttributionActivePay7DPerCount    int     `json:"attribution_active_pay_7d_per_count" bson:"attribution_active_pay_7_d_per_count"`
	Vote                              int     `json:"vote" bson:"vote"`
	LubanLivePayOrderStatCost         float64 `json:"luban_live_pay_order_stat_cost" bson:"luban_live_pay_order_stat_cost"`
	DownloadFinishRate                float64 `json:"download_finish_rate" bson:"download_finish_rate"`
	LoanCreditCost                    float64 `json:"loan_credit_cost" bson:"loan_credit_cost"`
	DeepConvertRate                   float64 `json:"deep_convert_rate" bson:"deep_convert_rate"`
	GameAddictionCost                 float64 `json:"game_addiction_cost" bson:"game_addiction_cost"`
	Shopping                          int     `json:"shopping" bson:"shopping"`
	LubanOrderCnt                     int     `json:"luban_order_cnt" bson:"luban_order_cnt"`
	AdvancedCreativePhoneClick        int     `json:"advanced_creative_phone_click" bson:"advanced_creative_phone_click"`
	LubanLiveSlidecartClickCnt        int     `json:"luban_live_slidecart_click_cnt" bson:"luban_live_slidecart_click_cnt"`
	Phone                             int     `json:"phone" bson:"phone"`
	WifiPlayRate                      float64 `json:"wifi_play_rate" bson:"wifi_play_rate"`
	Consult                           int     `json:"consult" bson:"consult"`
	WechatFirstPayCost                float64 `json:"wechat_first_pay_cost" bson:"wechat_first_pay_cost"`
	GameAddictionRate                 float64 `json:"game_addiction_rate" bson:"game_addiction_rate"`
	ActivePayRate                     float64 `json:"active_pay_rate" bson:"active_pay_rate"`
	DownloadStart                     int     `json:"download_start" bson:"download_start"`
	CardShow                          int     `json:"card_show" bson:"card_show"`
	MapSearch                         int     `json:"map_search" bson:"map_search"`
	InAppOrder                        int     `json:"in_app_order" bson:"in_app_order"`
	AttributionConvert                int64   `json:"attribution_convert" bson:"attribution_convert"`
	WechatFirstPayRate                float64 `json:"wechat_first_pay_rate" bson:"wechat_first_pay_rate"`
	AttributionGamePay7DCount         int     `json:"attribution_game_pay_7d_count" bson:"attribution_game_pay_7_d_count"`
	ClickDownload                     int     `json:"click_download" bson:"click_download"`
	PreLoanCredit                     int     `json:"pre_loan_credit" bson:"pre_loan_credit"`
	AttributionWechatLogin30DCost     float64 `json:"attribution_wechat_login_30d_cost" bson:"attribution_wechat_login_30_d_cost"`
	Share                             int     `json:"share" bson:"share"`
	ClickInstall                      int     `json:"click_install" bson:"click_install"`
	AdvancedCreativeFormSubmit        int     `json:"advanced_creative_form_submit" bson:"advanced_creative_form_submit"`
	DeepConvert                       int64   `json:"deep_convert" bson:"deep_convert"`
	LiveFansClubJoinCnt               int     `json:"live_fans_club_join_cnt" bson:"live_fans_club_join_cnt"`
	PhoneConnect                      int     `json:"phone_connect" bson:"phone_connect"`
	Id                                int64   `json:"id" bson:"id"`
	IesMusicClick                     int     `json:"ies_music_click" bson:"ies_music_click"`
	Play100FeedBreak                  int     `json:"play_100_feed_break" bson:"play_100_feed_break"`
	CouponSinglePage                  int     `json:"coupon_single_page" bson:"coupon_single_page"`
	Form                              int     `json:"form" bson:"form"`
	ValidPlayCost                     float64 `json:"valid_play_cost" bson:"valid_play_cost"`
	LiveWatchOneMinuteCount           int     `json:"live_watch_one_minute_count" bson:"live_watch_one_minute_count"`
	LubanOrderStatAmount              float64 `json:"luban_order_stat_amount" bson:"luban_order_stat_amount"`
	DownloadStartRate                 float64 `json:"download_start_rate" bson:"download_start_rate"`
	AttributionNextDayOpenRate        float64 `json:"attribution_next_day_open_rate" bson:"attribution_next_day_open_rate"`
	DownloadFinish                    int     `json:"download_finish" bson:"download_finish"`
	StatDatetime                      string  `json:"stat_datetime" bson:"stat_datetime"`
	Inventory                         string  `json:"inventory" bson:"inventory"`
	AttributionWechatFirstPay30DRate  float64 `json:"attribution_wechat_first_pay_30d_rate" bson:"attribution_wechat_first_pay_30_d_rate"`
	AttributionWechatPay30DAmount     float64 `json:"attribution_wechat_pay_30d_amount" bson:"attribution_wechat_pay_30_d_amount"`
	AdvancedCreativeCounselClick      int     `json:"advanced_creative_counsel_click" bson:"advanced_creative_counsel_click"`
	PhoneEffective                    int     `json:"phone_effective" bson:"phone_effective"`
	InAppDetailUv                     int     `json:"in_app_detail_uv" bson:"in_app_detail_uv"`
	AveragePlayTimePerPlay            float64 `json:"average_play_time_per_play" bson:"average_play_time_per_play"`
	PhoneConfirm                      int     `json:"phone_confirm" bson:"phone_confirm"`
	LocationClick                     int     `json:"location_click" bson:"location_click"`
	GamePayCount                      int     `json:"game_pay_count" bson:"game_pay_count"`
	ClickShopwindow                   int     `json:"click_shopwindow" bson:"click_shopwindow"`
	Register                          int     `json:"register" bson:"register"`
	InAppPay                          int     `json:"in_app_pay" bson:"in_app_pay"`
	LubanLiveFollowCnt                int     `json:"luban_live_follow_cnt" bson:"luban_live_follow_cnt"`
	NextDayOpenCost                   float64 `json:"next_day_open_cost" bson:"next_day_open_cost"`
	AdvancedCreativeFormClick         int     `json:"advanced_creative_form_click" bson:"advanced_creative_form_click"`
}
