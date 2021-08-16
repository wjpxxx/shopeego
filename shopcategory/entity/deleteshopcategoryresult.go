package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteShopCategoryResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	DeleteShopCategoryResponseResponseEntity	`json:"response"`
}
func (g DeleteShopCategoryResult) String() string {
    return lib.ObjectToString(g)
}
