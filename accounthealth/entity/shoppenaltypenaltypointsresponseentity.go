package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPenaltyPenaltyPointsResponseEntity struct{
    OverallPenaltyPoints	int	`json:"overall_penalty_points"`
    NonFulfillmentRate	int	`json:"non_fulfillment_rate"`
    LateShipmentRate	int	`json:"late_shipment_rate"`
    ListingViolations	int	`json:"listing_violations"`
    Others	int	`json:"others"`
}
func (g ShopPenaltyPenaltyPointsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
