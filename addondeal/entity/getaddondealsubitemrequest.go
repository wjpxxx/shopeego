package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetAddOnDealSubItemRequest struct{
    AddOnDealId	int64	`json:"add_on_deal_id"`
}
func (g GetAddOnDealSubItemRequest) String() string {
    return lib.ObjectToString(g)
}
