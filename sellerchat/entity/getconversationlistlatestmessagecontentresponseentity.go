package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetConversationListLatestMessageContentResponseEntity struct{
    Text	string	`json:"text"`
}
func (g GetConversationListLatestMessageContentResponseEntity) String() string {
    return lib.ObjectToString(g)
}
