package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type EndAddOnDealRequest struct{
    AddOnDealId	int64	`json:"add_on_deal_id"`
}
func (g EndAddOnDealRequest) String() string {
    return lib.ObjectToString(g)
}
