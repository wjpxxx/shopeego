package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GenerateFirstMileTrackingNumberResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	GenerateFirstMileTrackingNumberResponseResponseEntity	`json:"response"`
}
func (g GenerateFirstMileTrackingNumberResult) String() string {
    return lib.ObjectToString(g)
}
