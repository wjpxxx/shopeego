package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetBundleDealListRequest struct{
    PageSize	int	`json:"page_size"`
    TimeStatus	int	`json:"time_status"`
    PageNo	int	`json:"page_no"`
}
func (g GetBundleDealListRequest) String() string {
    return lib.ObjectToString(g)
}
