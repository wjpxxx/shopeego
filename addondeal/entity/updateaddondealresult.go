package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateAddOnDealResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	UpdateAddOnDealResponseResponseEntity	`json:"response"`
}
func (g UpdateAddOnDealResult) String() string {
    return lib.ObjectToString(g)
}
