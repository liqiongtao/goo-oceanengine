package goo_oceanengine

import "encoding/json"

type AdvertiserSelectResult struct {
	Code int `json:"code"`
	Data struct {
		List []AdvertiserSelectItem `json:"list"`
	} `json:"data"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
}

func (p AdvertiserSelectResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type AdvertiserSelectItem struct {
	AdvertiserId   int64  `json:"advertiser_id,omitempty"`
	AdvertiserName string `json:"advertiser_name"`
}
