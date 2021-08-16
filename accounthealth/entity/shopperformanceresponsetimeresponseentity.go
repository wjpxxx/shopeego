package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPerformanceResponseTimeResponseEntity struct{
    TotalData	ShopPerformanceTotalDataResponseEntity	`json:"total_data"`
}
func (g ShopPerformanceResponseTimeResponseEntity) String() string {
    return lib.ObjectToString(g)
}
