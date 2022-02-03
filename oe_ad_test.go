package goo_oceanengine

import (
	"fmt"
	"testing"
)

func TestAd_Get(t *testing.T) {
	accessToken := "8fade7017205ad92ce757475ac65b70f1609b3d2"

	params := AdGetParams{
		AdvertiserId: "1718542606637064",
		Fields:       nil,
		Page:         1,
		PageSize:     100,
	}

	rst := AD(Config{
		AppId:  "1719050127541279",
		Secret: "bd3873f97aa0910e4ae93378fbbae4e1bf691148",
	}).Get(params, accessToken)

	fmt.Println(rst)
}
