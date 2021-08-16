package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddAddOnDealMainItemResponseResponseEntity struct{
    MainItemList	[]AddAddOnDealMainItemMainItemListResponseEntity	`json:"main_item_list"`
    AddOnDealId	int64	`json:"add_on_deal_id"`
}
func (g AddAddOnDealMainItemResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
