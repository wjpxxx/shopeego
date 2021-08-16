package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPerformanceNonFulfillmentRateResponseEntity struct{
    TotalData	ShopPerformanceTotalDataResponseEntity	`json:"total_data"`
    CancellationRateData	ShopPerformanceCancellationRateDataResponseEntity	`json:"cancellation_rate_data"`
    ReturnRefundRateData	ShopPerformanceReturnRefundRateDataResponseEntity	`json:"return_refund_rate_data"`
}
func (g ShopPerformanceNonFulfillmentRateResponseEntity) String() string {
    return lib.ObjectToString(g)
}
