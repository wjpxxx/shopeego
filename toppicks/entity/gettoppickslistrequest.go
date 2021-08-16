package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetTopPicksListRequest struct{
}
func (g GetTopPicksListRequest) String() string {
    return lib.ObjectToString(g)
}
