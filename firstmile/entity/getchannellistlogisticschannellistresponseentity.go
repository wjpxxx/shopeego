package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetChannelListLogisticsChannelListResponseEntity struct{
    LogisticsChannelId	int64	`json:"logistics_channel_id"`
    LogisticsChannelName	string	`json:"logistics_channel_name"`
    ShipmentMethod	string	`json:"shipment_method"`
}
func (g GetChannelListLogisticsChannelListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
