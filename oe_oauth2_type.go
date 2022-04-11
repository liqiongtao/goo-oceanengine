package goo_oceanengine

type AccessTokenResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		AccessToken           string `json:"access_token"`
		ExpiresIn             int    `json:"expires_in"`
		RefreshToken          string `json:"refresh_token"`
		AdvertiserId          int    `json:"advertiser_id"`
		AdvertiserIds         []int  `json:"advertiser_ids"`
		RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
	} `json:"data"`
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

type AdvertiserGetResult struct {
	Code int `json:"code"`
	Data struct {
		List []struct {
			AccountRole    string        `json:"account_role"`
			AdvertiserId   int64         `json:"advertiser_id"`
			AdvertiserName string        `json:"advertiser_name"`
			AdvertiserRole int           `json:"advertiser_role"`
			Companist      []interface{} `json:"companist"`
			IsValid        bool          `json:"is_valid"`
		} `json:"list"`
	} `json:"data"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
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
