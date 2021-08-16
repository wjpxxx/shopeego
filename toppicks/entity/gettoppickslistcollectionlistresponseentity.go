package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetTopPicksListCollectionListResponseEntity struct{
    IsActivated	bool	`json:"is_activated"`
    ItemList	[]GetTopPicksListItemListResponseEntity	`json:"item_list"`
    TopPicksId	int64	`json:"top_picks_id"`
    Name	string	`json:"name"`
}
func (g GetTopPicksListCollectionListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
