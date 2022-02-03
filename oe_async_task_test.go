package goo_oceanengine

import (
	"fmt"
	"testing"
)

func TestAsyncTask_AsyncTaskCreate(t *testing.T) {
	accessToken := "8fade7017205ad92ce757475ac65b70f1609b3d2"

	params := AsyncTaskCreateParams{
		AdvertiserId: 1718542606637064,
		TaskName:     "",
		Force:        "",
		TaskType:     "",
		TaskParams: struct {
			StartDate  string   `json:"start_date,omitempty"`
			EndDate    string   `json:"end_date,omitempty"`
			GroupBy    string   `json:"group_by,omitempty"`
			OrderField string   `json:"order_field,omitempty"`
			OrderType  string   `json:"order_type,omitempty"`
			Fields     []string `json:"fields,omitempty"`
			Filtering  struct {
				CampaignId            []int64  `json:"campaign_id,omitempty"`
				AdId                  []int64  `json:"ad_id,omitempty"`
				CreativeId            []int64  `json:"creative_id,omitempty"`
				LandingType           []string `json:"landing_type,omitempty"`
				Pricing               []string `json:"pricing,omitempty"`
				InventoryType         []string `json:"inventory_type,omitempty"`
				CreativeMaterialModes []string `json:"creative_material_modes,omitempty"`
				ConvertType           []string `json:"convert_type,omitempty"`
				Platform              []string `json:"platform,omitempty"`
				Bidword               []string `json:"bidword,omitempty"`
				Query                 []string `json:"query,omitempty"`
				ImageMode             []string `json:"image_mode,omitempty"`
			} `json:"filtering,omitempty"`
		}{},
	}

	rst := AsyncTask(Config{
		AppId:  "1719050127541279",
		Secret: "bd3873f97aa0910e4ae93378fbbae4e1bf691148",
	}).Create(params, accessToken)

	fmt.Println(rst)
}

func TestAsyncTask_AsyncTaskGet(t *testing.T) {
	accessToken := "8fade7017205ad92ce757475ac65b70f1609b3d2"

	params := AsyncTaskGetParams{
		AdvertiserId: 1718542606637064,
		Filtering: struct {
			TaskIds  []int64 `json:"task_ids,omitempty"`
			TaskName string  `json:"task_name,omitempty"`
		}{},
		Page:     0,
		PageSize: 0,
	}

	rst := AsyncTask(Config{
		AppId:  "1719050127541279",
		Secret: "bd3873f97aa0910e4ae93378fbbae4e1bf691148",
	}).Get(params, accessToken)

	fmt.Println(rst)
}

func TestAsyncTask_AsyncTaskDownload(t *testing.T) {
	accessToken := "8fade7017205ad92ce757475ac65b70f1609b3d2"

	params := AsyncTaskDownloadParams{
		AdvertiserId: 1718542606637064,
		TaskId:       0,
		RangeFrom:    0,
		RangeTo:      0,
	}

	rst := AsyncTask(Config{
		AppId:  "1719050127541279",
		Secret: "bd3873f97aa0910e4ae93378fbbae4e1bf691148",
	}).Download(params, accessToken)

	fmt.Println(rst)
}
