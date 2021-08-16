package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type EndAddOnDealResponseResponseEntity struct{
    AddOnDealId	int64	`json:"add_on_deal_id"`
}
func (g EndAddOnDealResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
