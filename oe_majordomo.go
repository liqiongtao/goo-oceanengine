package goo_oceanengine

import (
	"fmt"
	goo_http_request "github.com/liqiongtao/googo.io/goo-http-request"
)

func Majordomo(config Config) majordomo {
	return majordomo{oceanengine{config: config}}
}

type majordomo struct {
	oceanengine
}

// 获取纵横组织下资产账户列表
// https://open.oceanengine.com/labels/7/docs/1696710519607296
// 获取纵横组织下的广告主ID列表。
func (r majordomo) AdvertiserSelect(advertiserId int64, accessToken string) (rst AdvertiserSelectResult) {
	data := fmt.Sprintf(`{"advertiser_id":%d}`, advertiserId)
	opt := goo_http_request.HeaderOption("Access-Token", accessToken)

	rst = AdvertiserSelectResult{}
	if err := r.get(MAJORDOMO_ADVERTISER_SELECT_URL, []byte(data), &rst, opt); err != nil {
		rst = AdvertiserSelectResult{Code: 5001, Message: err.Error()}
	}
	return
}
