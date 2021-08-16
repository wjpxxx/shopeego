package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetShopsByPartnerAuthedShopListResponseEntity struct{
    Region	string	`json:"region"`
    ShopId	int64	`json:"shop_id"`
    AuthTime	int	`json:"auth_time"`
    ExpireTime	int	`json:"expire_time"`
    SipAffiShopList	[]GetShopsByPartnerSipAffiShopListResponseEntity	`json:"sip_affi_shop_list"`
}
func (g GetShopsByPartnerAuthedShopListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
