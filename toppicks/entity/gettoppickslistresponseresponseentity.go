package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetTopPicksListResponseResponseEntity struct{
    CollectionList	[]GetTopPicksListCollectionListResponseEntity	`json:"collection_list"`
}
func (g GetTopPicksListResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
