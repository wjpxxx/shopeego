package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteConversationResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	DeleteConversationResponseResponseEntity	`json:"response"`
}
func (g DeleteConversationResult) String() string {
    return lib.ObjectToString(g)
}


type DeleteConversationResponseResponseEntity struct{

}

func (g DeleteConversationResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}