package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateAddOnDealSubItemRequest struct{
    AddOnDealId	int64	`json:"add_on_deal_id"`
    SubItemList	[]UpdateAddOnDealSubItemSubItemListRequestEntity	`json:"sub_item_list"`
}
func (g UpdateAddOnDealSubItemRequest) String() string {
    return lib.ObjectToString(g)
}
