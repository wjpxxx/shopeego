package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddTopPicksCollectionListResponseEntity struct{
    IsActivated	bool	`json:"is_activated"`
    ItemList	[]AddTopPicksItemListResponseEntity	`json:"item_list"`
    TopPicksId	int64	`json:"top_picks_id"`
    Name	string	`json:"name"`
}
func (g AddTopPicksCollectionListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
