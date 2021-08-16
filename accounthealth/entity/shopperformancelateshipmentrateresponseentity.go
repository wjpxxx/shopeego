package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPerformanceLateShipmentRateResponseEntity struct{
    TotalData	ShopPerformanceTotalDataResponseEntity	`json:"total_data"`
}
func (g ShopPerformanceLateShipmentRateResponseEntity) String() string {
    return lib.ObjectToString(g)
}
