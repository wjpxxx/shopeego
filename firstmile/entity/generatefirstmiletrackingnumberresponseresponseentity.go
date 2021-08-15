package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GenerateFirstMileTrackingNumberResponseResponseEntity struct{
    FirstMileTrackingNumberList	[]string	`json:"first_mile_tracking_number_list"`
}
func (g GenerateFirstMileTrackingNumberResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
