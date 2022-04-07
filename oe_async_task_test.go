package goo_oceanengine

import (
	"fmt"
	"testing"
)

func TestAsyncTask_AsyncTaskCreate(t *testing.T) {
	accessToken := ""

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
		AppId:  "",
		Secret: "",
	}).Create(params, accessToken)

	fmt.Println(rst)
}

func TestAsyncTask_AsyncTaskGet(t *testing.T) {
	accessToken := ""

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
		AppId:  "",
		Secret: "",
	}).Get(params, accessToken)

	fmt.Println(rst)
}

func TestAsyncTask_AsyncTaskDownload(t *testing.T) {
	accessToken := ""

	params := AsyncTaskDownloadParams{
		AdvertiserId: 1718542606637064,
		TaskId:       0,
		RangeFrom:    0,
		RangeTo:      0,
	}

	rst := AsyncTask(Config{
		AppId:  "",
		Secret: "",
	}).Download(params, accessToken)

	fmt.Println(rst)
}
