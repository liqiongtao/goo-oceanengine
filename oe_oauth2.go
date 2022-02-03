package goo_oceanengine

import (
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
// https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710502753295#item-link-%E8%AE%BE%E7%BD%AE%E6%8E%88%E6%9D%83URL
// app_id： 即应用ID，开发者创建应用的唯一标识，可通过应用管理查看；
// state：即自定义参数，可用于传递自定义信息，回调时会原样带回。常见应用比如将广告主账号作为自定义参数，回调时以区分授权码对应的广告主，其他应用方式可按实际需要选择
// scope：即授权范围，授权的权限范围，不传时代表当前应用拥有的所有权限。注意，权限范围只能在此应用拥有的权限范围之内。格式例如：scope=[1, 2, 3, 41]。具体权限范围取值见scope权限说明，每个权限后面的数字记为该权限的表示数值
// redirect_uri：即回调链接，由开发者自行提供和定义，用于授权成功跳转并接受回调信息。注意：redirect_uri 需要与APP应用的回调链接保持一致，否则会报错,回调地址示范：https://yourdomain.com/oauth2/callback/
func (oa oauth2) AuditUrl(redirectUri string) string {
	state := goo_utils.NonceStr8()
	scope := ""
	return fmt.Sprintf(OAUTH_URL, oa.config.AppId, state, scope, url.PathEscape(redirectUri))
}

// 获取Access Token
// https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710505596940#item-link-%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80
// Access-Token是调用接口时，操作指定广告账户的身份凭证，有效期为24小时
// Refresh-Token用于生成新access_token和refresh_token并且刷新时效达到续期的目的
func (oa oauth2) AccessToken(authCode string) goo_utils.Params {
	p := goo_utils.NewParams()
	p.Set("app_id", oa.config.AppId)
	p.Set("secret", oa.config.Secret)
	p.Set("grant_type", "auth_code")
	p.Set("auth_code", authCode)
	return oa.post(ACCESS_TOKEN_URL, p.JSON())
}

// 刷新Refresh Token
// https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710506097679#item-link-%E5%BA%94%E7%AD%94%E5%AD%97%E6%AE%B5
func (oa oauth2) RefreshToken(refreshToken string) goo_utils.Params {
	p := goo_utils.NewParams()
	p.Set("app_id", oa.config.AppId)
	p.Set("secret", oa.config.Secret)
	p.Set("grant_type", "refresh_token")
	p.Set("refresh_token", refreshToken)
	return oa.post(REFRESH_TOKEN_URL, p.JSON())
}

// 获取已授权账户
// https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710506097679#item-link-%E5%BA%94%E7%AD%94%E5%AD%97%E6%AE%B5
func (oa oauth2) AdvertiserGet(accessToken string) goo_utils.Params {
	p := goo_utils.NewParams()
	p.Set("app_id", oa.config.AppId)
	p.Set("secret", oa.config.Secret)
	p.Set("access_token", accessToken)
	return oa.get(ADVERTISER_GET_URL, p.JSON())
}

// 获取授权User信息
// https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710507039756#item-link-%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80
func (oa oauth2) UserInfo(accessToken string) goo_utils.Params {
	return oa.get(USER_INFO_URL, []byte{}, goo_http_request.HeaderOption("Access-Token", accessToken))
}
