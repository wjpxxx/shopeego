package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPerformanceFulfillmentResponseEntity struct{
    NonFulfillmentRate	ShopPerformanceNonFulfillmentRateResponseEntity	`json:"non_fulfillment_rate"`
    PreparationTime	ShopPerformancePreparationTimeResponseEntity	`json:"preparation_time"`
    LateShipmentRate	ShopPerformanceLateShipmentRateResponseEntity	`json:"late_shipment_rate"`
}
func (g ShopPerformanceFulfillmentResponseEntity) String() string {
    return lib.ObjectToString(g)
}
