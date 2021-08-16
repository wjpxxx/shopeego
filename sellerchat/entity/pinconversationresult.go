package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type PinConversationResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	PinConversationResponseResponseEntity	`json:"response"`
}
func (g PinConversationResult) String() string {
    return lib.ObjectToString(g)
}

type PinConversationResponseResponseEntity struct{

}

func (g PinConversationResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}