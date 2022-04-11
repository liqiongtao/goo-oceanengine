package goo_oceanengine

import (
	"fmt"
	"testing"
)

func TestCampaign_CampaignGet(t *testing.T) {
	accessToken := ""

	params := CampaignGetParams{
		AdvertiserId: 0,
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
		AppId:  "",
		Secret: "",
	}).Get(params, accessToken)

	fmt.Println(rst)
}
