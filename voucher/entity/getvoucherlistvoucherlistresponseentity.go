package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetVoucherListVoucherListResponseEntity struct{
    VoucherId	int64	`json:"voucher_id"`
    VoucherCode	string	`json:"voucher_code"`
    VoucherName	string	`json:"voucher_name"`
    VoucherType	int	`json:"voucher_type"`
    RewardType	int	`json:"reward_type"`
    UsageQuantity	int	`json:"usage_quantity"`
    CurrentUsage	int	`json:"current_usage"`
    StartTime	int	`json:"start_time"`
    EndTime	int	`json:"end_time"`
    IsAdmin	bool	`json:"is_admin"`
    VoucherPurpose	int	`json:"voucher_purpose"`
    DiscountAmount	float32	`json:"discount_amount"`
    Percentage	float32	`json:"percentage"`
    CmtVoucherStatus	int	`json:"cmt_voucher_status"`
}
func (g GetVoucherListVoucherListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
