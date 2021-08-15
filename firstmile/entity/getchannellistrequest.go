package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetChannelListRequest struct{
    Region	string	`json:"region"`
}
func (g GetChannelListRequest) String() string {
    return lib.ObjectToString(g)
}
