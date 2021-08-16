package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateShopCategoryResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	UpdateShopCategoryResponseResponseEntity	`json:"response"`
}
func (g UpdateShopCategoryResult) String() string {
    return lib.ObjectToString(g)
}
