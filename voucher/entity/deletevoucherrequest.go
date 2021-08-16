package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteVoucherRequest struct{
    VoucherId	int64	`json:"voucher_id"`
}
func (g DeleteVoucherRequest) String() string {
    return lib.ObjectToString(g)
}
