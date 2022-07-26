package goo_oceanengine

import "encoding/json"

type AdvertiserSelectResult struct {
	Code int `json:"code" bson:"code"`
	Data struct {
		List []AdvertiserSelectItem `json:"list" bson:"list"`
	} `json:"data" bson:"data"`
	Message   string `json:"message" bson:"message"`
	RequestId string `json:"request_id" bson:"request_id"`
}

func (p AdvertiserSelectResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type AdvertiserSelectItem struct {
	AdvertiserId   int64  `json:"advertiser_id,omitempty"`
	AdvertiserName string `json:"advertiser_name"`
}
