package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetConversationListPageResultResponseEntity struct{
    PageSize	int	`json:"page_size"`
    NextCursor	GetConversationListNextCursorResponseEntity	`json:"next_cursor"`
    More	bool	`json:"more"`
}
func (g GetConversationListPageResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}
