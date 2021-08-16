package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPerformanceOtherListingViolationsResponseEntity struct{
    TotalData	ShopPerformanceTotalDataResponseEntity	`json:"total_data"`
}
func (g ShopPerformanceOtherListingViolationsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
