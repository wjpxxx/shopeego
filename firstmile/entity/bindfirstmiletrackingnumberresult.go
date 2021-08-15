package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type BindFirstMileTrackingNumberResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Warning	[]BindFirstMileTrackingNumberWarningResponseEntity	`json:"warning"`
    Response	BindFirstMileTrackingNumberResponseResponseEntity	`json:"response"`
}
func (g BindFirstMileTrackingNumberResult) String() string {
    return lib.ObjectToString(g)
}
