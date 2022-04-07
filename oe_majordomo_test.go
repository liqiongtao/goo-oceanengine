package goo_oceanengine

import (
	"fmt"
	"testing"
)

func TestMajordomo_AdvertiserSelect(t *testing.T) {
	accessToken := ""

	rst := Majordomo(Config{
		AppId:  "",
		Secret: "",
	}).AdvertiserSelect(1718542607453191, accessToken)

	fmt.Println(rst)
}
