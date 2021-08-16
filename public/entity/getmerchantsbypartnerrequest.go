package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetMerchantsByPartnerRequest struct{
    PageSize	int	`json:"page_size"`
    PageNo	int	`json:"page_no"`
}
func (g GetMerchantsByPartnerRequest) String() string {
    return lib.ObjectToString(g)
}
