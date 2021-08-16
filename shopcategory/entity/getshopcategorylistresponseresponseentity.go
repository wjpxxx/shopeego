package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetShopCategoryListResponseResponseEntity struct{
    ShopCategoryList	[]GetShopCategoryListShopCategoryListResponseEntity	`json:"shop_category_list"`
    TotalCount	int	`json:"total_count"`
}
func (g GetShopCategoryListResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
