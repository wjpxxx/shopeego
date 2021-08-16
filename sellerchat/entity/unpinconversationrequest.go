package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UnpinConversationRequest struct{
    ConversationId	int64	`json:"conversation_id"`
}
func (g UnpinConversationRequest) String() string {
    return lib.ObjectToString(g)
}
