package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPerformanceSevereListingViolationsResponseEntity struct{
    TotalData	ShopPerformanceTotalDataResponseEntity	`json:"total_data"`
    SpamListingData	ShopPerformanceSpamListingDataResponseEntity	`json:"spam_listing_data"`
    CounterfeitIpInfringementData	ShopPerformanceCounterfeitIpInfringementDataResponseEntity	`json:"counterfeit_ip_infringement_data"`
    ProhibitedListingData	ShopPerformanceProhibitedListingDataResponseEntity	`json:"prohibited_listing_data"`
}
func (g ShopPerformanceSevereListingViolationsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
