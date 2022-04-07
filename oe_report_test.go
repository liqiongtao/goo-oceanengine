package goo_oceanengine

import (
	"fmt"
	"testing"
	"time"
)

func TestReport_ReportAdvertiserGet(t *testing.T) {
	accessToken := ""

	nw := time.Now()

	params := ReportAdvertiserGetParams{
		AdvertiserId: "1718542606637064",
		StartDate:    nw.Format("2006-01-02"),
		EndDate:      nw.Format("2006-01-02"),
		Page:         1,
		PageSize:     100,
	}

	rst := Report(Config{
		AppId:  "",
		Secret: "",
	}).AdvertiserGet(params, accessToken)

	fmt.Println(rst)
}

func TestReport_ReportCampaignGet(t *testing.T) {
	accessToken := ""

	nw := time.Now()

	params := ReportCampaignGetParams{
		AdvertiserId: "1718542606637064",
		StartDate:    nw.Format("2006-01-02"),
		EndDate:      nw.Format("2006-01-02"),
		GroupBy:      []string{"STAT_GROUP_BY_FIELD_ID", "STAT_GROUP_BY_FIELD_STAT_TIME"},
		Page:         1,
		PageSize:     100,
	}

	rst := Report(Config{
		AppId:  "",
		Secret: "",
	}).CampaignGet(params, accessToken)

	fmt.Println(rst)
}

func TestReport_ReportAdGet(t *testing.T) {
	accessToken := ""

	nw := time.Now()

	params := ReportAdGetParams{
		AdvertiserId:    "1718542606637064",
		StartDate:       nw.Format("2006-01-02"),
		EndDate:         nw.Format("2006-01-02"),
		GroupBy:         []string{"STAT_GROUP_BY_FIELD_ID", "STAT_GROUP_BY_FIELD_STAT_TIME"},
		TimeGranularity: "STAT_TIME_GRANULARITY_HOURLY",
		Page:            1,
		PageSize:        100,
	}

	rst := Report(Config{
		AppId:  "",
		Secret: "",
	}).AdGet(params, accessToken)

	for _, i := range rst.Get("data.list").Array() {
		fmt.Println(
			i.Get("campaign_id").Int64(),
			i.Get("ad_id").Int64(),
			i.Get("stat_datetime").String(),
			i.Get("click").Int64(),
			i.Get("show").Int64(),
			i.Get("convert").Int64(),
			i.Get("cost").Float64(),
			i.Get("convert_cost").Float64(),
			i.Get("convert_rate").Float64(),
			i.Get("ctr").Float64(),
		)
	}

	//fmt.Println(rst)
}

func TestReport_ReportCreativeGet(t *testing.T) {
	accessToken := ""

	nw := time.Now()

	params := ReportCreativeGetParams{
		AdvertiserId: "1718542606637064",
		StartDate:    nw.Format("2006-01-02"),
		EndDate:      nw.Format("2006-01-02"),
		GroupBy:      []string{"STAT_GROUP_BY_FIELD_ID", "STAT_GROUP_BY_FIELD_STAT_TIME"},
		Page:         1,
		PageSize:     100,
	}

	rst := Report(Config{
		AppId:  "",
		Secret: "",
	}).CreativeGet(params, accessToken)

	fmt.Println(rst)
}
