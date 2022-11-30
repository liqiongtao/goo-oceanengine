package goo_oceanengine

import (
	"encoding/json"
	"fmt"
	goo_http_request "github.com/liqiongtao/googo.io/goo-http-request"
	goo_utils "github.com/liqiongtao/googo.io/goo-utils"
	"net/url"
)

func OAuth2(config Config) oauth2 {
	return oauth2{oceanengine{config: config}}
}

type oauth2 struct {
	oceanengine
}

// 授权URL
// https://open.oceanengine.com/labels/7/docs/1696710502753295
// app_id： 即应用ID，开发者创建应用的唯一标识，可通过应用管理查看；
// state：即自定义参数，可用于传递自定义信息，回调时会原样带回。常见应用比如将广告主账号作为自定义参数，回调时以区分授权码对应的广告主，其他应用方式可按实际需要选择
// scope：即授权范围，授权的权限范围，不传时代表当前应用拥有的所有权限。注意，权限范围只能在此应用拥有的权限范围之内。格式例如：scope=[1, 2, 3, 41]。具体权限范围取值见scope权限说明，每个权限后面的数字记为该权限的表示数值
// redirect_uri：即回调链接，由开发者自行提供和定义，用于授权成功跳转并接受回调信息。注意：redirect_uri 需要与APP应用的回调链接保持一致，否则会报错,回调地址示范：https://yourdomain.com/oauth2/callback/
func (oa oauth2) AuditUrl(redirectUri string, args ...string) string {
	b, _ := json.Marshal(&args)
	scope := ""
	return fmt.Sprintf(OAUTH_URL, oa.config.AppId, url.PathEscape(string(b)), scope, url.PathEscape(redirectUri))
}

// 获取Access Token
// https://open.oceanengine.com/labels/7/docs/1696710505596940
// Access-Token是调用接口时，操作指定广告账户的身份凭证，有效期为24小时
// Refresh-Token用于生成新access_token和refresh_token并且刷新时效达到续期的目的
func (oa oauth2) AccessToken(authCode string) (rst AccessTokenResult) {
	m := goo_utils.M{
		"app_id":     oa.config.AppId,
		"secret":     oa.config.Secret,
		"grant_type": "auth_code",
		"auth_code":  authCode,
	}

	rst = AccessTokenResult{}
	if err := oa.post(ACCESS_TOKEN_URL, m.Json(), &rst); err != nil {
		rst = AccessTokenResult{Code: 5001, Message: err.Error()}
	}
	return
}

// 刷新Refresh Token
// https://open.oceanengine.com/labels/7/docs/1696710506097679
func (oa oauth2) RefreshToken(refreshToken string) (rst RefreshTokenResult) {
	m := goo_utils.M{
		"app_id":        oa.config.AppId,
		"secret":        oa.config.Secret,
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	}

	rst = RefreshTokenResult{}
	if err := oa.post(REFRESH_TOKEN_URL, m.Json(), &rst); err != nil {
		rst = RefreshTokenResult{Code: 5001, Message: err.Error()}
	}
	return
}

// 获取已授权账户
// https://open.oceanengine.com/labels/7/docs/1696710506574848
func (oa oauth2) AdvertiserGet(accessToken string) (rst AdvertiserGetResult) {
	m := goo_utils.M{
		"app_id":       oa.config.AppId,
		"secret":       oa.config.Secret,
		"access_token": accessToken,
	}

	rst = AdvertiserGetResult{}
	if err := oa.get(ADVERTISER_GET_URL, m.Json(), &rst); err != nil {
		rst = AdvertiserGetResult{Code: 5001, Message: err.Error()}
	}
	return
}

// 获取授权User信息
// https://open.oceanengine.com/labels/7/docs/1696710507039756
func (oa oauth2) UserInfo(accessToken string) (rst UserInfoResult) {
	opt := goo_http_request.HeaderOption("Access-Token", accessToken)

	rst = UserInfoResult{}
	if err := oa.get(USER_INFO_URL, []byte{}, &rst, opt); err != nil {
		rst = UserInfoResult{Code: 5001, Message: err.Error()}
	}
	return
}
