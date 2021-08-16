package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOfferToggleStatusResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	GetOfferToggleStatusResponseResponseEntity	`json:"response"`
}
func (g GetOfferToggleStatusResult) String() string {
    return lib.ObjectToString(g)
}
