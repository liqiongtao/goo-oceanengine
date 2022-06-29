package goo_oceanengine

import (
	"fmt"
	"testing"
)

func TestDmp_DataSourceFileUpload(t *testing.T) {
	accessToken := ""

	params := DataSourceFileUploadParams{
		AdvertiserId: 0,
		File:         "",
	}

	rst := DMP(Config{
		AppId:  "",
		Secret: "",
		Debug:  true,
	}).DataSourceFileUpload(params, accessToken)

	fmt.Println(rst)
}
