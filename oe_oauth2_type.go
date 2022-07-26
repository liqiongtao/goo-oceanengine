package goo_oceanengine

import "encoding/json"

type AccessTokenResult struct {
	Code    int    `json:"code" bson:"code"`
	Message string `json:"message" bson:"message"`
	Data    struct {
		AccessToken           string  `json:"access_token" bson:"access_token"`
		ExpiresIn             int     `json:"expires_in" bson:"expires_in"`
		RefreshToken          string  `json:"refresh_token" bson:"refresh_token"`
		RefreshTokenExpiresIn int     `json:"refresh_token_expires_in" bson:"refresh_token_expires_in"`
		AdvertiserId          int64   `json:"advertiser_id" bson:"advertiser_id"`
		AdvertiserName        string  `json:"advertiser_name" bson:"advertiser_name"`
		AdvertiserIds         []int64 `json:"advertiser_ids" bson:"advertiser_ids"`
	} `json:"data" bson:"data"`
}

func (p AccessTokenResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type RefreshTokenResult struct {
	Code    int    `json:"code" bson:"code"`
	Message string `json:"message" bson:"message"`
	Data    struct {
		AccessToken           string `json:"access_token" bson:"access_token"`
		ExpiresIn             int    `json:"expires_in" bson:"expires_in"`
		RefreshToken          string `json:"refresh_token" bson:"refresh_token"`
		RefreshTokenExpiresIn int    `json:"refresh_token_expires_in" bson:"refresh_token_expires_in"`
	} `json:"data" bson:"data"`
}

func (p RefreshTokenResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type AdvertiserGetResult struct {
	Code int `json:"code" bson:"code"`
	Data struct {
		List []AdvertiserInfo `json:"list" bson:"list"`
	} `json:"data" bson:"data"`
	Message   string `json:"message" bson:"message"`
	RequestId string `json:"request_id" bson:"request_id"`
}

func (p AdvertiserGetResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type AdvertiserInfo struct {
	AccountRole    string `json:"account_role"`
	AdvertiserId   int64  `json:"advertiser_id"`
	AdvertiserName string `json:"advertiser_name"`
	AdvertiserRole int    `json:"advertiser_role"`
	IsValid        bool   `json:"is_valid"`
}

type UserInfoResult struct {
	Message string `json:"message" bson:"message"`
	Code    int    `json:"code" bson:"code"`
	Data    struct {
		Id          int    `json:"id" bson:"id"`
		Email       string `json:"email" bson:"email"`
		DisplayName string `json:"display_name" bson:"display_name"`
	} `json:"data" bson:"data"`
	RequestId string `json:"request_id" bson:"request_id"`
}

func (p UserInfoResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}
