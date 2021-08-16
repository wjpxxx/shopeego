package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPerformancePreparationTimeResponseEntity struct{
    TotalData	ShopPerformanceTotalDataResponseEntity	`json:"total_data"`
}
func (g ShopPerformancePreparationTimeResponseEntity) String() string {
    return lib.ObjectToString(g)
}
