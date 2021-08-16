package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddVoucherRequest struct{
    VoucherName	string	`json:"voucher_name"`
    VoucherCode	string	`json:"voucher_code"`
    StartTime	int	`json:"start_time"`
    EndTime	int	`json:"end_time"`
    VoucherType	int	`json:"voucher_type"`
    RewardType	int	`json:"reward_type"`
    UsageQuantity	int	`json:"usage_quantity"`
    MinBasketPrice	float32	`json:"min_basket_price"`
    DiscountAmount	float32	`json:"discount_amount"`
    Percentage	int	`json:"percentage"`
    MaxPrice	float32	`json:"max_price"`
    DisplayChannelList	[]int	`json:"display_channel_list"`
    ItemIdList	[]int64	`json:"item_id_list"`
}
func (g AddVoucherRequest) String() string {
    return lib.ObjectToString(g)
}
