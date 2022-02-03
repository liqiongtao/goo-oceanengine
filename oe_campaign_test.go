package goo_oceanengine

import (
	"fmt"
	"testing"
)

func TestCampaign_CampaignGet(t *testing.T) {
	accessToken := "8fade7017205ad92ce757475ac65b70f1609b3d2"

	params := CampaignGetParams{
		AdvertiserId: "1718542606637064",
		filtering: struct {
			Ids                []int64 `json:"ids,omitempty"`
			CampaignName       string  `json:"campaign_name,omitempty"`
			LandingType        string  `json:"landing_type,omitempty"`
			Status             string  `json:"status,omitempty"`
			CampaignCreateTime string  `json:"campaign_create_time,omitempty"`
		}{},
		Fields:   nil,
		Page:     0,
		PageSize: 0,
	}

	rst := Campaign(Config{
		AppId:  "1719050127541279",
		Secret: "bd3873f97aa0910e4ae93378fbbae4e1bf691148",
	}).Get(params, accessToken)

	fmt.Println(rst)
}
