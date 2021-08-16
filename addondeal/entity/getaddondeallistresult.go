package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetAddOnDealListResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	GetAddOnDealListResponseResponseEntity	`json:"response"`
}
func (g GetAddOnDealListResult) String() string {
    return lib.ObjectToString(g)
}
