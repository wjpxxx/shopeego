package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//GetShopInfoResult
type GetShopInfoResult struct {
	ShopName     string               `json:"shop_name"`
	Region       string               `json:"region"`
	Status       string               `json:"status"`
	SipAffiShops SipAffiShopsList `json:"sip_affi_shops"`
	IsCb         bool                 `json:"is_cb"`
	IsCnsc       bool                 `json:"is_cnsc"`
	commonentity.Result
	AuthTime   int `json:"auth_time"`
	ExpireTime int `json:"expire_time"`
}

//String
func (g GetShopInfoResult) String() string {
	return lib.ObjectToString(g)
}


type SipAffiShopsList []SipAffiShopsEntity

//String
func (g SipAffiShopsList) String() string {
	return lib.ObjectToString(g)
}