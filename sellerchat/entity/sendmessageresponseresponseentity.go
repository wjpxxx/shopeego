package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SendMessageResponseResponseEntity struct{
    Content	SendMessageContentResponseEntity	`json:"content"`
    Region	string	`json:"region"`
    ToId	int64	`json:"to_id"`
    CreatedTimestamp	int	`json:"created_timestamp"`
    ConversationId	string	`json:"conversation_id"`
    MessageType	string	`json:"message_type"`
    MessageId	string	`json:"message_id"`
    MessageOption	int	`json:"message_option"`
}
func (g SendMessageResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
