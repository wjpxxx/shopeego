package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetShopCategoryListShopCategoryListResponseEntity struct{
    ShopCategoryId	int64	`json:"shop_category_id"`
    Status	string	`json:"status"`
    Name	string	`json:"name"`
    SortWeight	int	`json:"sort_weight"`
}
func (g GetShopCategoryListShopCategoryListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
