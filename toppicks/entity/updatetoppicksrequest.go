package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateTopPicksRequest struct{
    TopPicksId	int64	`json:"top_picks_id"`
    Name	string	`json:"name"`
    ItemIdList	[]int64	`json:"item_id_list"`
    IsActivated	bool	`json:"is_activated"`
}
func (g UpdateTopPicksRequest) String() string {
    return lib.ObjectToString(g)
}
