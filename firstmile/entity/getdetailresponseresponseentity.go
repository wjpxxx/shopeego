package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetDetailResponseResponseEntity struct{
    LogisticsChannelId	int64	`json:"logistics_channel_id"`
    FirstMileTrackingNumber	string	`json:"first_mile_tracking_number"`
    ShipmentMethod	string	`json:"shipment_method"`
    Status	string	`json:"status"`
    DeclareDate	string	`json:"declare_date"`
    More	bool	`json:"more"`
    OrderList	[]GetDetailOrderListResponseEntity	`json:"order_list"`
    NextCursor	string	`json:"next_cursor"`
}
func (g GetDetailResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
