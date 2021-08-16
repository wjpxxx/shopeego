package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SendMessageContentResponseEntity struct{
    Text	string	`json:"text"`
}
func (g SendMessageContentResponseEntity) String() string {
    return lib.ObjectToString(g)
}
