package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UnpinConversationResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	UnpinConversationResponseResponseEntity	`json:"response"`
}
func (g UnpinConversationResult) String() string {
    return lib.ObjectToString(g)
}


type UnpinConversationResponseResponseEntity struct{

}

func (g UnpinConversationResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}