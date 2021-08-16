package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteAddOnDealSubItemResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	DeleteAddOnDealSubItemResponseResponseEntity	`json:"response"`
}
func (g DeleteAddOnDealSubItemResult) String() string {
    return lib.ObjectToString(g)
}
