package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPerformanceListingViolationsResponseEntity struct{
    SevereListingViolations	ShopPerformanceSevereListingViolationsResponseEntity	`json:"severe_listing_violations"`
    PreOrderListing	ShopPerformancePreOrderListingResponseEntity	`json:"pre_order_listing"`
    OtherListingViolations	ShopPerformanceOtherListingViolationsResponseEntity	`json:"other_listing_violations"`
}
func (g ShopPerformanceListingViolationsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
