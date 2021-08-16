package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPerformanceResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    OverallPerformance	int	`json:"overall_performance"`
    ListingViolations	ShopPerformanceListingViolationsResponseEntity	`json:"listing_violations"`
    Fulfillment	ShopPerformanceFulfillmentResponseEntity	`json:"fulfillment"`
    CustomerService	ShopPerformanceCustomerServiceResponseEntity	`json:"customer_service"`
    CustomerSatisfaction	ShopPerformanceCustomerSatisfactionResponseEntity	`json:"customer_satisfaction"`
}
func (g ShopPerformanceResult) String() string {
    return lib.ObjectToString(g)
}
