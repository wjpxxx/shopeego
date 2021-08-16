package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddShopCategoryRequest struct{
    Name	string	`json:"name"`
    SortWeight	int	`json:"sort_weight"`
}
func (g AddShopCategoryRequest) String() string {
    return lib.ObjectToString(g)
}
