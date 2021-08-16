package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ReadConversationResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	ReadConversationResponseResponseEntity	`json:"response"`
}
func (g ReadConversationResult) String() string {
    return lib.ObjectToString(g)
}

type ReadConversationResponseResponseEntity struct{

}

func (g ReadConversationResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}