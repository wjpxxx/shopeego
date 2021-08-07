package shop

import (
	"github.com/wjpxxx/letgo/lib"
	shopeeConfig "github.com/wjpxxx/shopeego/config"
	shopEntity "github.com/wjpxxx/shopeego/shop/entity"
)

//Shop
type Shop struct {
	Config *shopeeConfig.Config
}

//GetShopInfo
//@Title Use this call to get information of shop
//@Description https://open.shopee.com/documents?module=92&type=1&id=536&version=2
func (m *Shop) GetShopInfo() shopEntity.GetShopInfoResult {
	method := "shop/get_shop_info"
	result := shopEntity.GetShopInfoResult{}
	params := lib.InRow{}
	err := m.Config.HttpGet(method, params, &result)
	if err != nil {
		result.Error = err.Error()
	}
	return result
}

//GetProfile
//@Title This API support to get information of shop.
//@Description https://open.shopee.com/documents?module=92&type=1&id=584&version=2
func (m *Shop) GetProfile() shopEntity.GetProfileResult {
	method := "shop/get_profile"
	result := shopEntity.GetProfileResult{}
	params := lib.InRow{}
	err := m.Config.HttpGet(method, params, &result)
	if err != nil {
		result.Error = err.Error()
	}
	return result
}

//UpdateProfile
//@Title This API support to let sellers to update the shop name, shop logo, and shop description.
//@Description https://open.shopee.com/documents?module=92&type=1&id=585&version=2
func (m *Shop) UpdateProfile(shopName, shopLogo, description string) shopEntity.UpdateProfileResult {
	method := "shop/update_profile"
	result := shopEntity.UpdateProfileResult{}
	params := lib.InRow{
		"shop_name":   shopName,
		"shop_logo":   shopLogo,
		"description": description,
	}
	err := m.Config.HttpPost(method, params, &result)
	if err != nil {
		result.Error = err.Error()
	}
	return result
}
