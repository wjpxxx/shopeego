package public

import (
	"github.com/wjpxxx/letgo/lib"
	shopeeConfig "github.com/wjpxxx/shopeego/config"
	publicentity "github.com/wjpxxx/shopeego/public/entity"
)

//Public
type Public struct {
	Config *shopeeConfig.Config
}

//GetShopsByPartner
//@Title get basic info of shops which have authorized to the partner.
//@Description https://open.shopee.com/documents?module=99&type=1&id=663&version=2
func (m *Public) GetShopsByPartner(params publicentity.GetShopsByPartnerRequest) publicentity.GetShopsByPartnerResult {
    method := "public/get_shops_by_partner"
    result := publicentity.GetShopsByPartnerResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//GetMerchantsByPartner
//@Title Use this API to get basic info of merchants which have authorized to the partner.
//@Description https://open.shopee.com/documents?module=99&type=1&id=664&version=2
func (m *Public) GetMerchantsByPartner(params publicentity.GetMerchantsByPartnerRequest) publicentity.GetMerchantsByPartnerResult {
    method := "public/get_merchants_by_partner"
    result := publicentity.GetMerchantsByPartnerResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
