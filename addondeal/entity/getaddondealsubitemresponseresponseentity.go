package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetAddOnDealSubItemResponseResponseEntity struct{
    SubItemList	[]GetAddOnDealSubItemSubItemListResponseEntity	`json:"sub_item_list"`
    AddOnDealId	int64	`json:"add_on_deal_id"`
}
func (g GetAddOnDealSubItemResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
