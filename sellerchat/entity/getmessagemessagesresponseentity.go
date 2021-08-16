package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetMessageMessagesResponseEntity struct{
    MessageId	string	`json:"message_id"`
    MessageType	string	`json:"message_type"`
    FromId	int64	`json:"from_id"`
    FromShopId	int64	`json:"from_shop_id"`
    ToId	int64	`json:"to_id"`
    ToShopId	int64	`json:"to_shop_id"`
    ConversationId	string	`json:"conversation_id"`
    CreatedTimestamp	int	`json:"created_timestamp"`
    Region	string	`json:"region"`
    Status	string	`json:"status"`
    Source	string	`json:"source"`
    Content	GetMessageContentResponseEntity	`json:"content"`
    MessageOption	int	`json:"message_option"`
}
func (g GetMessageMessagesResponseEntity) String() string {
    return lib.ObjectToString(g)
}
