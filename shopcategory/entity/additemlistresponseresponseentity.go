package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddItemListResponseResponseEntity struct{
    InvalidItemIdList	[]AddItemListInvalidItemIdListResponseEntity	`json:"invalid_item_id_list"`
    ShopCategoryId	int64	`json:"shop_category_id"`
    CurrentCount	int	`json:"current_count"`
}
func (g AddItemListResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
