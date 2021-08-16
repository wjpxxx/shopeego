package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetAddOnDealResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	GetAddOnDealResponseResponseEntity	`json:"response"`
}
func (g GetAddOnDealResult) String() string {
    return lib.ObjectToString(g)
}
