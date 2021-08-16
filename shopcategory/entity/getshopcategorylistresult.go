package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetShopCategoryListResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	GetShopCategoryListResponseResponseEntity	`json:"response"`
}
func (g GetShopCategoryListResult) String() string {
    return lib.ObjectToString(g)
}
