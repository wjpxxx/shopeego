package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//GetAccessTokenResult
type GetAccessTokenResult struct {
	commonentity.Result
	commonentity.ShopInfo
	MerchantIDList []int `json:"merchant_id_list"`
	ShopIDList     []int `json:"shop_id_list"`
}

//String
func (g GetAccessTokenResult) String() string {
	return lib.ObjectToString(g)
}

//RefreshAccessTokenResult
type RefreshAccessTokenResult struct {
	commonentity.Result
	commonentity.ShopInfo
	PartnerID  int64 `json:"partner_id"`
	MerchantID int64 `json:"merchant_id"`
}

//String
func (r RefreshAccessTokenResult) String() string {
	return lib.ObjectToString(r)
}
