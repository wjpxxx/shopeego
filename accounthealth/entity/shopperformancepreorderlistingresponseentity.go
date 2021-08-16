package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPerformancePreOrderListingResponseEntity struct{
    PercentPreOrderListingData	ShopPerformancePercentPreOrderListingDataResponseEntity	`json:"percent_pre_order_listing_data"`
    NumberOfDaysPreOrderListingExceedsTargetDa	ShopPerformanceNumberOfDaysPreOrderListingExceedsTargetDaResponseEntity	`json:"number_of_days_pre_order_listing_exceeds_target_da"`
}
func (g ShopPerformancePreOrderListingResponseEntity) String() string {
    return lib.ObjectToString(g)
}
