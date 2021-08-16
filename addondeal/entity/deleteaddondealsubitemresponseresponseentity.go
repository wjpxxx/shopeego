package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteAddOnDealSubItemResponseResponseEntity struct{
    SubItemList	[]DeleteAddOnDealSubItemSubItemListResponseEntity	`json:"sub_item_list"`
    AddOnDealId	int64	`json:"add_on_deal_id"`
}
func (g DeleteAddOnDealSubItemResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
