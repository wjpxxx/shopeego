package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddAddOnDealSubItemResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	AddAddOnDealSubItemResponseResponseEntity	`json:"response"`
}
func (g AddAddOnDealSubItemResult) String() string {
    return lib.ObjectToString(g)
}
