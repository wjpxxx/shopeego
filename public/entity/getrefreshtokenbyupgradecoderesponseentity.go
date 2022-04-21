package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetRefreshTokenByUpgradeCodeResponseEntity struct{
    SuccessShopIdList	[]int64	`json:"success_shop_id_list"`
    RefreshToken	string	`json:"refresh_token"`
    FailureList	[]FailureList	`json:"failure_list"`
}
func (g GetRefreshTokenByUpgradeCodeResponseEntity) String() string {
    return lib.ObjectToString(g)
}

type FailureList struct{
	ShopId int64 `json:"shop_id"`
	FailedReason string `json:"failed_reason"`
}
func (g FailureList) String() string {
    return lib.ObjectToString(g)
}