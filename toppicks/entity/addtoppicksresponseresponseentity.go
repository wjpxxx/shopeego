package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddTopPicksResponseResponseEntity struct{
    CollectionList	[]AddTopPicksCollectionListResponseEntity	`json:"collection_list"`
}
func (g AddTopPicksResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
