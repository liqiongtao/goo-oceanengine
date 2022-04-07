package goo_oceanengine

import (
	"fmt"
	goo_log "github.com/liqiongtao/googo.io/goo-log"
	"testing"
)

func TestOauth2_AuditUrl(t *testing.T) {
	url := OAuth2(Config{
		AppId:  "",
		Secret: "",
	}).AuditUrl("http://api.abc.com/oauth2/callback")

	goo_log.Debug(url)
}

func TestOauth2_RefreshToken(t *testing.T) {
	refreshToken := ""

	rst := OAuth2(Config{
		AppId:  "",
		Secret: "",
	}).RefreshToken(refreshToken)

	fmt.Println(rst)
}

func TestOauth2_AdvertiserGet(t *testing.T) {
	accessToken := ""

	rst := OAuth2(Config{
		AppId:  "",
		Secret: "",
	}).AdvertiserGet(accessToken)

	fmt.Println(rst)
}

func TestOauth2_UserInfo(t *testing.T) {
	accessToken := ""

	rst := OAuth2(Config{
		AppId:  "",
		Secret: "",
	}).UserInfo(accessToken)

	fmt.Println(rst)
}
