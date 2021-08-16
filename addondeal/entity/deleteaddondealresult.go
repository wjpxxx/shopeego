package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteAddOnDealResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	DeleteAddOnDealResponseResponseEntity	`json:"response"`
}
func (g DeleteAddOnDealResult) String() string {
    return lib.ObjectToString(g)
}
