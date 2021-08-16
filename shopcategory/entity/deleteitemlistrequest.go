package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteItemListRequest struct{
    ShopCategoryId	int64	`json:"shop_category_id"`
    Items	[]int	`json:"items"`
}
func (g DeleteItemListRequest) String() string {
    return lib.ObjectToString(g)
}
