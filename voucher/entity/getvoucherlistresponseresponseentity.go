package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetVoucherListResponseResponseEntity struct{
    More	bool	`json:"more"`
    VoucherList	[]GetVoucherListVoucherListResponseEntity	`json:"voucher_list"`
}
func (g GetVoucherListResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
