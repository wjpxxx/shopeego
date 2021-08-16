package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetAddOnDealListResponseResponseEntity struct{
    AddOnDealList	[]GetAddOnDealListAddOnDealListResponseEntity	`json:"add_on_deal_list"`
    More	bool	`json:"more"`
}
func (g GetAddOnDealListResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
