package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddShopCategoryResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	AddShopCategoryResponseResponseEntity	`json:"response"`
}
func (g AddShopCategoryResult) String() string {
    return lib.ObjectToString(g)
}
