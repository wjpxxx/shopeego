package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddShopCategoryResponseResponseEntity struct{
    ShopCategoryId	int64	`json:"shop_category_id"`
}
func (g AddShopCategoryResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
