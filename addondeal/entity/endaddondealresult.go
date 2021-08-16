package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type EndAddOnDealResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	EndAddOnDealResponseResponseEntity	`json:"response"`
}
func (g EndAddOnDealResult) String() string {
    return lib.ObjectToString(g)
}
