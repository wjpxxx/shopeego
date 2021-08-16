package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteShopCategoryResponseResponseEntity struct{
    ShopCategoryId	int64	`json:"shop_category_id"`
    Msg	string	`json:"msg"`
}
func (g DeleteShopCategoryResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
