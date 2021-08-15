package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetTrackingNumberListFirstMileTrackingNumberListResponseEntity struct{
    FirstMileTrackingNumber	string	`json:"first_mile_tracking_number"`
    Status	string	`json:"status"`
    DeclareDate	string	`json:"declare_date"`
}
func (g GetTrackingNumberListFirstMileTrackingNumberListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
