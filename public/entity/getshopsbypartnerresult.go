package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetShopsByPartnerResult struct{
    AuthedShopList	[]GetShopsByPartnerAuthedShopListResponseEntity	`json:"authed_shop_list"`
    RequestId	string	`json:"request_id"`
    More	bool	`json:"more"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
}
func (g GetShopsByPartnerResult) String() string {
    return lib.ObjectToString(g)
}
