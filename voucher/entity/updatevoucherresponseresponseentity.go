package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateVoucherResponseResponseEntity struct{
    VoucherId	int64	`json:"voucher_id"`
}
func (g UpdateVoucherResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
