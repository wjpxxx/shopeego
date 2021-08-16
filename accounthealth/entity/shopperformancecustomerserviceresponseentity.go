package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPerformanceCustomerServiceResponseEntity struct{
    ResponseRate	ShopPerformanceResponseRateResponseEntity	`json:"response_rate"`
    ResponseTime	ShopPerformanceResponseTimeResponseEntity	`json:"response_time"`
}
func (g ShopPerformanceCustomerServiceResponseEntity) String() string {
    return lib.ObjectToString(g)
}
