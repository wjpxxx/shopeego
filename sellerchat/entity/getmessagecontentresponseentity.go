package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetMessageContentResponseEntity struct{
    Text	string	`json:"text"`
}
func (g GetMessageContentResponseEntity) String() string {
    return lib.ObjectToString(g)
}
