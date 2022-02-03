package goo_oceanengine

import (
	"fmt"
	goo_http_request "github.com/liqiongtao/googo.io/goo-http-request"
	goo_utils "github.com/liqiongtao/googo.io/goo-utils"
)

func Majordomo(config Config) majordomo {
	return majordomo{oceanengine{config: config}}
}

type majordomo struct {
	oceanengine
}

// 获取纵横组织下资产账户列表
// https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710519607296#item-link-%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80
// 获取纵横组织下的广告主ID列表。
func (r majordomo) AdvertiserSelect(advertiserId int64, accessToken string) goo_utils.Params {
	data := fmt.Sprintf(`{"advertiser_id":%d}`, advertiserId)
	return r.get(MAJORDOMO_ADVERTISER_SELECT_URL, []byte(data), goo_http_request.HeaderOption("Access-Token", accessToken))
}
