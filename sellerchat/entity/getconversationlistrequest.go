package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetConversationListRequest struct{
    Direction	string	`json:"direction"`
    Type	string	`json:"type"`
    NextTimestampNano	int	`json:"next_timestamp_nano"`
    PageSize	int	`json:"page_size"`
}
func (g GetConversationListRequest) String() string {
    return lib.ObjectToString(g)
}
