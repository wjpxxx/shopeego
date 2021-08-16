package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetAddOnDealSubItemResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	GetAddOnDealSubItemResponseResponseEntity	`json:"response"`
}
func (g GetAddOnDealSubItemResult) String() string {
    return lib.ObjectToString(g)
}
