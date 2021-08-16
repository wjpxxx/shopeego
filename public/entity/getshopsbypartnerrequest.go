package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetShopsByPartnerRequest struct{
    PageSize	int	`json:"page_size"`
    PageNo	int	`json:"page_no"`
}
func (g GetShopsByPartnerRequest) String() string {
    return lib.ObjectToString(g)
}
