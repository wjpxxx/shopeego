package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddAddOnDealResponseResponseEntity struct{
    AddOnDealId	int64	`json:"add_on_deal_id"`
}
func (g AddAddOnDealResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
