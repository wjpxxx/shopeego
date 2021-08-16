package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteAddOnDealRequest struct{
    AddOnDealId	int64	`json:"add_on_deal_id"`
}
func (g DeleteAddOnDealRequest) String() string {
    return lib.ObjectToString(g)
}
