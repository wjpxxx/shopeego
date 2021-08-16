package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UnreadConversationRequest struct{
    ConversationId	int64	`json:"conversation_id"`
}
func (g UnreadConversationRequest) String() string {
    return lib.ObjectToString(g)
}
