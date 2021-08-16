package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SendMessageResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	SendMessageResponseResponseEntity	`json:"response"`
}
func (g SendMessageResult) String() string {
    return lib.ObjectToString(g)
}
