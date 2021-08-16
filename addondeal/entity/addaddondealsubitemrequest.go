package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddAddOnDealSubItemRequest struct{
    AddOnDealId	int64	`json:"add_on_deal_id"`
    SubItemList	[]AddAddOnDealSubItemSubItemListRequestEntity	`json:"sub_item_list"`
}
func (g AddAddOnDealSubItemRequest) String() string {
    return lib.ObjectToString(g)
}
