package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UnreadConversationResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	UnreadConversationResponseResponseEntity	`json:"response"`
}
func (g UnreadConversationResult) String() string {
    return lib.ObjectToString(g)
}

type UnreadConversationResponseResponseEntity struct{

}

func (g UnreadConversationResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}