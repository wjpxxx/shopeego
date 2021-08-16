package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOneConversationResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	GetOneConversationResponseResponseEntity	`json:"response"`
}
func (g GetOneConversationResult) String() string {
    return lib.ObjectToString(g)
}
