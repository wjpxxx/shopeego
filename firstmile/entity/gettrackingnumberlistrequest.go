package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetTrackingNumberListRequest struct{
    FromDate	string	`json:"from_date"`
    ToDate	string	`json:"to_date"`
    PageSize	int	`json:"page_size"`
    Cursor	string	`json:"cursor"`
}
func (g GetTrackingNumberListRequest) String() string {
    return lib.ObjectToString(g)
}
