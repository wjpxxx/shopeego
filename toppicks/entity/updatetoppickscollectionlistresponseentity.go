package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateTopPicksCollectionListResponseEntity struct{
    IsActivated	bool	`json:"is_activated"`
    ItemList	[]UpdateTopPicksItemListResponseEntity	`json:"item_list"`
    TopPicksId	int64	`json:"top_picks_id"`
    Name	string	`json:"name"`
}
func (g UpdateTopPicksCollectionListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
