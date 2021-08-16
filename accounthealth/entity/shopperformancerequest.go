package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPerformanceRequest struct{
}
func (g ShopPerformanceRequest) String() string {
    return lib.ObjectToString(g)
}
