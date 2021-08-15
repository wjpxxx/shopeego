package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GenerateFirstMileTrackingNumberSellerInfoRequestEntity struct{
    Address	string	`json:"address"`
    Name	string	`json:"name"`
    Zipcode	string	`json:"zipcode"`
    Region	string	`json:"region"`
    Phone	string	`json:"phone"`
}
func (g GenerateFirstMileTrackingNumberSellerInfoRequestEntity) String() string {
    return lib.ObjectToString(g)
}
