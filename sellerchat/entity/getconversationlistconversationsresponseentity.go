package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetConversationListConversationsResponseEntity struct{
    ConversationId	string	`json:"conversation_id"`
    ToId	int64	`json:"to_id"`
    ToName	string	`json:"to_name"`
    ToAvatar	string	`json:"to_avatar"`
    ShopId	int64	`json:"shop_id"`
    UnreadCount	int	`json:"unread_count"`
    Pinned	bool	`json:"pinned"`
    LastReadMessageId	string	`json:"last_read_message_id"`
    LatestMessageId	string	`json:"latest_message_id"`
    LatestMessageType	string	`json:"latest_message_type"`
    LatestMessageContent	GetConversationListLatestMessageContentResponseEntity	`json:"latest_message_content"`
    LatestMessageFromId	int64	`json:"latest_message_from_id"`
    LastMessageTimestamp	int	`json:"last_message_timestamp"`
    LastMessageOption	int	`json:"last_message_option"`
    MaxGeneralOptionHideTime	string	`json:"max_general_option_hide_time"`
}
func (g GetConversationListConversationsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
