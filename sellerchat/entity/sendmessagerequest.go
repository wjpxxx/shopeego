package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SendMessageRequest struct{
    ToId	int64	`json:"to_id"`
    MessageType	string	`json:"message_type"`
    Content	SendMessageContentRequestEntity	`json:"content"`
}
func (g SendMessageRequest) String() string {
    return lib.ObjectToString(g)
}
