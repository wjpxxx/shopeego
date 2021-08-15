package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UnbindFirstMileTrackingNumberResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Warning	[]UnbindFirstMileTrackingNumberWarningResponseEntity	`json:"warning"`
    Response	UnbindFirstMileTrackingNumberResponseResponseEntity	`json:"response"`
}
func (g UnbindFirstMileTrackingNumberResult) String() string {
    return lib.ObjectToString(g)
}
