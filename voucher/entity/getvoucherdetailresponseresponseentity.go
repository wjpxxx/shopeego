package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetVoucherDetailResponseResponseEntity struct{
    VoucherId	int64	`json:"voucher_id"`
    VoucherCode	string	`json:"voucher_code"`
    VoucherName	string	`json:"voucher_name"`
    VoucherType	int	`json:"voucher_type"`
    RewardType	int	`json:"reward_type"`
    UsageQuantity	int	`json:"usage_quantity"`
    CurrentUsage	int	`json:"current_usage"`
    StartTime	int	`json:"start_time"`
    EndTime		int `json:"end_time"`
    IsAdmin	bool	`json:"is_admin"`
    VoucherPurpose	int	`json:"voucher_purpose"`
    DisplayChannelList	[]int	`json:"display_channel_list"`
    MinBasketPrice	float32	`json:"min_basket_price"`
    Percentage	int	`json:"percentage"`
    MaxPrice	float32	`json:"max_price"`
    DiscountAmount	float32	`json:"discount_amount"`
    CmtVoucherStatus	int	`json:"cmt_voucher_status"`
    ItemIdList	[]int64	`json:"item_id_list"`
}
func (g GetVoucherDetailResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
