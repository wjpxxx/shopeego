package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateAddOnDealSubItemResponseResponseEntity struct{
    SubItemList	[]UpdateAddOnDealSubItemSubItemListResponseEntity	`json:"sub_item_list"`
}
func (g UpdateAddOnDealSubItemResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
