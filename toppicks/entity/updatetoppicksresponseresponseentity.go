package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateTopPicksResponseResponseEntity struct{
    CollectionList	[]UpdateTopPicksCollectionListResponseEntity	`json:"collection_list"`
}
func (g UpdateTopPicksResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
