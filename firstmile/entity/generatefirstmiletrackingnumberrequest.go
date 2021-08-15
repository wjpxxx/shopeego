package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GenerateFirstMileTrackingNumberRequest struct{
    DeclareDate	string	`json:"declare_date"`
    Quantity	int	`json:"quantity"`
    SellerInfo	GenerateFirstMileTrackingNumberSellerInfoRequestEntity	`json:"seller_info"`
}
func (g GenerateFirstMileTrackingNumberRequest) String() string {
    return lib.ObjectToString(g)
}
