package goo_oceanengine

import (
	"fmt"
	"testing"
)

func TestAd_Get(t *testing.T) {
	accessToken := ""

	params := AdGetParams{
		AdvertiserId: 0,
		Fields:       nil,
		Page:         1,
		PageSize:     100,
	}

	rst := AD(Config{
		AppId:  "",
		Secret: "",
	}).Get(params, accessToken)

	fmt.Println(rst)
}
