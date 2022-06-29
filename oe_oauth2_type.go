package goo_oceanengine

import "encoding/json"

type AccessTokenResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		AccessToken           string  `json:"access_token"`
		ExpiresIn             int     `json:"expires_in"`
		RefreshToken          string  `json:"refresh_token"`
		RefreshTokenExpiresIn int     `json:"refresh_token_expires_in"`
		AdvertiserId          int64   `json:"advertiser_id"`
		AdvertiserName        string  `json:"advertiser_name"`
		AdvertiserIds         []int64 `json:"advertiser_ids"`
	} `json:"data"`
}

func (p AccessTokenResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type RefreshTokenResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		AccessToken           string `json:"access_token"`
		ExpiresIn             int    `json:"expires_in"`
		RefreshToken          string `json:"refresh_token"`
		RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
	} `json:"data"`
}

func (p RefreshTokenResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}

type AdvertiserGetResult struct {
	Code int `json:"code"`
	Data struct {
		List []AdvertiserInfo `json:"list"`
	} `json:"data"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
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
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		Id          int    `json:"id"`
		Email       string `json:"email"`
		DisplayName string `json:"display_name"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

func (p UserInfoResult) Json() []byte {
	b, _ := json.Marshal(&p)
	return b
}
