package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddAddOnDealMainItemRequest struct{
    AddOnDealId	int64	`json:"add_on_deal_id"`
    MainItemList	[]AddAddOnDealMainItemMainItemListRequestEntity	`json:"main_item_list"`
}
func (g AddAddOnDealMainItemRequest) String() string {
    return lib.ObjectToString(g)
}
