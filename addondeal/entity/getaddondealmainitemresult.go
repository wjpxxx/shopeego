package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetAddOnDealMainItemResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	GetAddOnDealMainItemResponseResponseEntity	`json:"response"`
}
func (g GetAddOnDealMainItemResult) String() string {
    return lib.ObjectToString(g)
}
