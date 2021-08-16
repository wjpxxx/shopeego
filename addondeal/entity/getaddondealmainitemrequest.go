package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetAddOnDealMainItemRequest struct{
    AddOnDealId	int64	`json:"add_on_deal_id"`
}
func (g GetAddOnDealMainItemRequest) String() string {
    return lib.ObjectToString(g)
}
