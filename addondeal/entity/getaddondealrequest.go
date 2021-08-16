package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetAddOnDealRequest struct{
    AddOnDealId	int64	`json:"add_on_deal_id"`
}
func (g GetAddOnDealRequest) String() string {
    return lib.ObjectToString(g)
}
