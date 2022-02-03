package goo_oceanengine

import (
	"fmt"
	"testing"
)

func TestMajordomo_AdvertiserSelect(t *testing.T) {
	accessToken := "8fade7017205ad92ce757475ac65b70f1609b3d2"

	rst := Majordomo(Config{
		AppId:  "1719050127541279",
		Secret: "bd3873f97aa0910e4ae93378fbbae4e1bf691148",
	}).AdvertiserSelect(1718542607453191, accessToken)

	fmt.Println(rst)
}
