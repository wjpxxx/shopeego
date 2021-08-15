package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetEscrowListRequest struct{
    ReleaseTimeFrom	int	`json:"release_time_from"`
    ReleaseTimeTo	int	`json:"release_time_to"`
    PageSize	int	`json:"page_size"`
    PageNo	int	`json:"page_no"`
}
func (g GetEscrowListRequest) String() string {
    return lib.ObjectToString(g)
}
