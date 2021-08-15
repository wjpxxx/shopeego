package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetDetailOrderListResponseEntity struct{
    OrderSn	string	`json:"order_sn"`
    PackageNumber	string	`json:"package_number"`
    SlsTrackingNumber	string	`json:"sls_tracking_number"`
    PickUpDone	bool	`json:"pick_up_done"`
    ArrivedTransitWarehouse	bool	`json:"arrived_transit_warehouse"`
}
func (g GetDetailOrderListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
