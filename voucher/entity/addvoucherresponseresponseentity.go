package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddVoucherResponseResponseEntity struct{
    VoucherId	int64	`json:"voucher_id"`
}
func (g AddVoucherResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
