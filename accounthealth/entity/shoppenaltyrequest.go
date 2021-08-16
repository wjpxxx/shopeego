package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPenaltyRequest struct{
}
func (g ShopPenaltyRequest) String() string {
    return lib.ObjectToString(g)
}
