package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetTrackingNumberListResponseResponseEntity struct{
    More	bool	`json:"more"`
    FirstMileTrackingNumberList	[]GetTrackingNumberListFirstMileTrackingNumberListResponseEntity	`json:"first_mile_tracking_number_list"`
    NextCursor	string	`json:"next_cursor"`
}
func (g GetTrackingNumberListResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
