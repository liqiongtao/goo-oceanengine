package goo_oceanengine

type AdvertiserSelectResult struct {
	Code int `json:"code"`
	Data struct {
		List []AdvertiserSelectItem `json:"list"`
	} `json:"data"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
}

type AdvertiserSelectItem struct {
	AdvertiserId   int64  `json:"advertiser_id,omitempty"`
	AdvertiserName string `json:"advertiser_name"`
}
