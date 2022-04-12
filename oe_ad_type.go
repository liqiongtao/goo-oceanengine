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
	} `json:"filtering,omitempty"`                 // 过滤条件，若此字段不传，或传空则视为无限制条件
	Fields   []string `json:"fields,omitempty"`    // 查询字段集合, 如果指定, 则返回结果数组中，不支持audience下字段筛选
	Page     int32    `json:"page,omitempty"`      // 页数默认值: 1，page范围为[1,99999]
	PageSize int32    `json:"page_size,omitempty"` // 页面大小默认值:10，page_size范围为[1,100]
}

func (p AdGetParams) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type AdGetResult struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
	Data      struct {
		List []struct {
			AdModifyTime               string        `json:"ad_modify_time"`
			ParamsType                 string        `json:"params_type"`
			AppThumbnails              []string      `json:"app_thumbnails"`
			ValueOptimizedOpen         int           `json:"value_optimized_open"`
			FeedDeliverySearch         string        `json:"feed_delivery_search"`
			DeepBidType                string        `json:"deep_bid_type"`
			SmartBidType               string        `json:"smart_bid_type"`
			ExternalActions            []string      `json:"external_actions"`
			BudgetMode                 string        `json:"budget_mode"`
			AdCreateTime               string        `json:"ad_create_time"`
			DownloadType               string        `json:"download_type"`
			OptStatus                  string        `json:"opt_status"`
			Name                       string        `json:"name"`
			ExternalUrl                interface{}   `json:"external_url"`
			LandingPageStayTime        int           `json:"landing_page_stay_time"`
			TrackUrlSendType           string        `json:"track_url_send_type"`
			InventoryCatalog           string        `json:"inventory_catalog"`
			SubscribeUrl               string        `json:"subscribe_url"`
			Pricing                    string        `json:"pricing"`
			Status                     string        `json:"status"`
			DeliveryRange              string        `json:"delivery_range"`
			HideIfConverted            string        `json:"hide_if_converted"`
			DpaRecommendType           int           `json:"dpa_recommend_type"`
			StartTime                  string        `json:"start_time"`
			LubanRoiGoal               float32       `json:"luban_roi_goal"`
			CategoryType               string        `json:"category_type"`
			DpaLbs                     int           `json:"dpa_lbs"`
			AdvertiserId               int64         `json:"advertiser_id"`
			UdShop                     interface{}   `json:"ud_shop"`
			AdId                       int64         `json:"ad_id"`
			ValueOptimizedType         int           `json:"value_optimized_type"`
			TrackUrl                   []string      `json:"track_url"`
			TargetCvr                  int64         `json:"target_cvr"`
			InheritType                string        `json:"inherit_type"`
			AppIntroduction            string        `json:"app_introduction"`
			StoreproPackId             int           `json:"storepro_pack_id"`
			Budget                     float64       `json:"budget"`
			DpaCategories              []int         `json:"dpa_categories"`
			AdvancedCreativeType       string        `json:"advanced_creative_type"`
			CpaBid                     float64       `json:"cpa_bid"`
			AwemeAccount               string        `json:"aweme_account"`
			DpaProductTarget           []interface{} `json:"dpa_product_target"`
			AudiencePackageId          interface{}   `json:"audience_package_id"`
			ExternalUrlParams          interface{}   `json:"external_url_params"`
			GamePackageThumbnailIds    interface{}   `json:"game_package_thumbnail_ids"`
			DpaProvince                interface{}   `json:"dpa_province"`
			IntelligentFlowSwitch      string        `json:"intelligent_flow_switch"`
			StoreproUnit               interface{}   `json:"storepro_unit"`
			VideoPlayTrackUrl          []string      `json:"video_play_track_url"`
			AppType                    string        `json:"app_type"`
			ModifyTime                 string        `json:"modify_time"`
			DpaExternalUrlField        interface{}   `json:"dpa_external_url_field"`
			CampaignId                 int64         `json:"campaign_id"`
			DpaOpenUrlType             interface{}   `json:"dpa_open_url_type"`
			PromotionType              interface{}   `json:"promotion_type"`
			FormIndex                  interface{}   `json:"form_index"`
			ConvertedTimeDuration      string        `json:"converted_time_duration"`
			AssetIds                   []int64       `json:"asset_ids"`
			OpenUrlParams              string        `json:"open_url_params"`
			DownloadUrl                string        `json:"download_url"`
			GamePackageBatchId         interface{}   `json:"game_package_batch_id"`
			ScheduleType               string        `json:"schedule_type"`
			DpaOpenUrls                interface{}   `json:"dpa_open_urls"`
			SceneInventory             string        `json:"scene_inventory"`
			ProductId                  interface{}   `json:"product_id"`
			AutoUpdateKeyword          interface{}   `json:"auto_update_keyword"`
			DeepExternalAction         string        `json:"deep_external_action"`
			UnionVideoType             interface{}   `json:"union_video_type"`
			Ulink                      interface{}   `json:"ulink"`
			VideoPlayEffectiveTrackUrl []interface{} `json:"video_play_effective_track_url"`
			ActionTrackUrl             []string      `json:"action_track_url"`
			DeepCpabid                 float64       `json:"deep_cpabid"`
			ConvertId                  int64         `json:"convert_id"`
			DpaExternalUrls            interface{}   `json:"dpa_external_urls"`
			AutoInheritSwitch          string        `json:"auto_inherit_switch"`
			Bid                        float64       `json:"bid"`
			AssetId                    int64         `json:"asset_id"`
			DownloadMode               string        `json:"download_mode"`
			DpaProducts                interface{}   `json:"dpa_products"`
			SmartInventory             string        `json:"smart_inventory"`
			RoiGoal                    float64       `json:"roi_goal"`
			TrackUrlGroupType          interface{}   `json:"track_url_group_type"`
			DpaCity                    interface{}   `json:"dpa_city"`
			InheritedAdvertiserId      interface{}   `json:"inherited_advertiser_id"`
			DpaAdtype                  interface{}   `json:"dpa_adtype"`
			GamePackageDesc            interface{}   `json:"game_package_desc"`
			HideIfExists               int           `json:"hide_if_exists"`
			VideoPlayDoneTrackUrl      []interface{} `json:"video_play_done_track_url"`
			EndTime                    string        `json:"end_time"`
			Audience                   struct {
				DpaRtaRecommendType interface{}   `json:"dpa_rta_recommend_type"`
				Ac                  []interface{} `json:"ac"`
				City                []int         `json:"city"`
				RegionVersion       string        `json:"region_version"`
				InterestWords       []interface{} `json:"interest_words"`
				Action              struct {
					ActionDays       interface{} `json:"action_days"`
					ActionWords      interface{} `json:"action_words"`
					ActionCategories interface{} `json:"action_categories"`
					ActionScene      interface{} `json:"action_scene"`
				} `json:"action"`
				InterestActionMode        string        `json:"interest_action_mode"`
				OwnAwemeNumber            interface{}   `json:"own_aweme_number"`
				DeviceType                interface{}   `json:"device_type"`
				BusinessIds               []interface{} `json:"business_ids"`
				DeviceBrand               []interface{} `json:"device_brand"`
				FilterAwemeAbnormalActive interface{}   `json:"filter_aweme_abnormal_active"`
				DpaRtaSwitch              interface{}   `json:"dpa_rta_switch"`
				Carrier                   []interface{} `json:"carrier"`
				Age                       []interface{} `json:"age"`
				UserType                  []interface{} `json:"user_type"`
				AutoExtendEnabled         int           `json:"auto_extend_enabled"`
				IosOsv                    interface{}   `json:"ios_osv"`
				Gender                    string        `json:"gender"`
				ArticleCategory           []interface{} `json:"article_category"`
				AwemeFansNumbers          []interface{} `json:"aweme_fans_numbers"`
				ExcludeFlowPackage        []interface{} `json:"exclude_flow_package"`
				ExcludeCustomActions      []interface{} `json:"exclude_custom_actions"`
				AwemeFanCategories        []interface{} `json:"aweme_fan_categories"`
				Platform                  []string      `json:"platform"`
				AwemeFanAccounts          []interface{} `json:"aweme_fan_accounts"`
				AwemeFanTimeScope         interface{}   `json:"aweme_fan_time_scope"`
				AppCategory               interface{}   `json:"app_category"`
				ActivateType              []interface{} `json:"activate_type"`
				AdTag                     []int         `json:"ad_tag"`
				Career                    interface{}   `json:"career"`
				LaunchPrice               []interface{} `json:"launch_price"`
				RetargetingTagsInclude    []interface{} `json:"retargeting_tags_include"`
				DpaLocalAudience          interface{}   `json:"dpa_local_audience"`
				AppBehaviorTarget         interface{}   `json:"app_behavior_target"`
				AppIds                    []interface{} `json:"app_ids"`
				FilterOwnAwemeFans        interface{}   `json:"filter_own_aweme_fans"`
				IncludeCustomActions      []interface{} `json:"include_custom_actions"`
				District                  string        `json:"district"`
				InterestTags              []interface{} `json:"interest_tags"`
				SuperiorPopularityType    string        `json:"superior_popularity_type"`
				AutoExtendTargets         []string      `json:"auto_extend_targets"`
				LocationType              string        `json:"location_type"`
				RetargetingTagsExclude    []interface{} `json:"retargeting_tags_exclude"`
				Geolocation               []interface{} `json:"geolocation"`
				AndroidOsv                interface{}   `json:"android_osv"`
				FlowPackage               []interface{} `json:"flow_package"`
				FilterAwemeFansCount      interface{}   `json:"filter_aweme_fans_count"`
				InterestCategories        []interface{} `json:"interest_categories"`
				AwemeFanBehaviors         []interface{} `json:"aweme_fan_behaviors"`
			} `json:"audience"`
			DpaOpenUrlField    interface{} `json:"dpa_open_url_field"`
			OpenUrl            string      `json:"open_url"`
			FlowControlMode    string      `json:"flow_control_mode"`
			QuickAppUrl        string      `json:"quick_app_url"`
			AuditRejectReason  interface{} `json:"audit_reject_reason"`
			Id                 int64       `json:"id"`
			Package            string      `json:"package"`
			AdGroupId          interface{} `json:"ad_group_id"`
			StoreType          string      `json:"store_type"`
			ExternalAction     string      `json:"external_action"`
			AppDesc            interface{} `json:"app_desc"`
			InventoryType      []string    `json:"inventory_type"`
			ProductPlatformId  int64       `json:"product_platform_id"`
			AdvertiserStoreIds []int64     `json:"advertiser_store_ids"`
			DeliveryPhase      string      `json:"delivery_phase"`
			AdjustCpa          int         `json:"adjust_cpa"`
			ScheduleTime       string      `json:"schedule_time"`
			FormId             int64       `json:"form_id"`
		} `json:"list"`
		PageInfo struct {
			Page        int `json:"page"`
			PageSize    int `json:"page_size"`
			TotalPage   int `json:"total_page"`
			TotalNumber int `json:"total_number"`
		} `json:"page_info"`
	} `json:"data"`
}
