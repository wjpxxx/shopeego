package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetChannelListResponseResponseEntity struct{
    LogisticsChannelList	[]GetChannelListLogisticsChannelListResponseEntity	`json:"logistics_channel_list"`
}
func (g GetChannelListResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
