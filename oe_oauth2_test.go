package goo_oceanengine

import (
	"fmt"
	goo_log "github.com/liqiongtao/googo.io/goo-log"
	"testing"
)

func TestOauth2_AuditUrl(t *testing.T) {
	url := OAuth2(Config{
		AppId:  "1719050127541279",
		Secret: "bd3873f97aa0910e4ae93378fbbae4e1bf691148",
	}).AuditUrl("http://api.abc.com/oauth2/callback")

	goo_log.Debug(url)
}

func TestOauth2_RefreshToken(t *testing.T) {
	refreshToken := "96d29c7857f072e7a8057a7c09b91776bcc2c1c6"

	rst := OAuth2(Config{
		AppId:  "1719050127541279",
		Secret: "bd3873f97aa0910e4ae93378fbbae4e1bf691148",
	}).RefreshToken(refreshToken)

	fmt.Println(rst)
}

func TestOauth2_AdvertiserGet(t *testing.T) {
	accessToken := "8fade7017205ad92ce757475ac65b70f1609b3d2"

	rst := OAuth2(Config{
		AppId:  "1719050127541279",
		Secret: "bd3873f97aa0910e4ae93378fbbae4e1bf691148",
	}).AdvertiserGet(accessToken)

	fmt.Println(rst)
}

func TestOauth2_UserInfo(t *testing.T) {
	accessToken := "8fade7017205ad92ce757475ac65b70f1609b3d2"

	rst := OAuth2(Config{
		AppId:  "1719050127541279",
		Secret: "bd3873f97aa0910e4ae93378fbbae4e1bf691148",
	}).UserInfo(accessToken)

	fmt.Println(rst)
}
