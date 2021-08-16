package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddItemListRequest struct{
    ShopCategoryId	int64	`json:"shop_category_id"`
    Items	[]int	`json:"items"`
}
func (g AddItemListRequest) String() string {
    return lib.ObjectToString(g)
}
