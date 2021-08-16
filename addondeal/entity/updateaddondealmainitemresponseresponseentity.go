package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateAddOnDealMainItemResponseResponseEntity struct{
    MainItemList	[]UpdateAddOnDealMainItemMainItemListResponseEntity	`json:"main_item_list"`
}
func (g UpdateAddOnDealMainItemResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
