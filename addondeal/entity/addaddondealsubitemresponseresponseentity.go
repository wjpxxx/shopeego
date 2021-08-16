package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddAddOnDealSubItemResponseResponseEntity struct{
    SubItemList	[]AddAddOnDealSubItemSubItemListResponseEntity	`json:"sub_item_list"`
    AddOnDealId	int64	`json:"add_on_deal_id"`
}
func (g AddAddOnDealSubItemResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
