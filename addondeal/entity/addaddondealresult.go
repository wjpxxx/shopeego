package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddAddOnDealResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	AddAddOnDealResponseResponseEntity	`json:"response"`
}
func (g AddAddOnDealResult) String() string {
    return lib.ObjectToString(g)
}
