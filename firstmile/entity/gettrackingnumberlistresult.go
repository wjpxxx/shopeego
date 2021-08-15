package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetTrackingNumberListResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	GetTrackingNumberListResponseResponseEntity	`json:"response"`
}
func (g GetTrackingNumberListResult) String() string {
    return lib.ObjectToString(g)
}
