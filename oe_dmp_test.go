package goo_oceanengine

import (
	"fmt"
	"testing"
)

func TestDMP(t *testing.T) {
	accessToken := ""

	params := CustomAudienceSelectParams{
		AdvertiserId: 0,
		SelectType:   0,
	}

	rst := DMP(Config{
		AppId:  "",
		Secret: "",
	}).CustomAudienceSelect(params, accessToken)

	fmt.Println(rst)
}
