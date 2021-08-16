package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateVoucherRequest struct{
    VoucherId	int64	`json:"voucher_id"`
    VoucherName	string	`json:"voucher_name"`
    StartTime	int	`json:"start_time"`
    EndTime	int	`json:"end_time"`
    UsageQuantity	int	`json:"usage_quantity"`
    MinBasketPrice	float32	`json:"min_basket_price"`
    DiscountAmount	float32	`json:"discount_amount"`
    Percentage	int	`json:"percentage"`
    MaxPrice	float32	`json:"max_price"`
    DisplayChannelList	[]int	`json:"display_channel_list"`
    ItemIdList	[]int64	`json:"item_id_list"`
}
func (g UpdateVoucherRequest) String() string {
    return lib.ObjectToString(g)
}
