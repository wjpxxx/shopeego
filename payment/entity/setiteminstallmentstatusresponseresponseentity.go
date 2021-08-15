package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetItemInstallmentStatusResponseResponseEntity struct{
    ItemInstallmentList	[]SetItemInstallmentStatusItemInstallmentListResponseEntity	`json:"item_installment_list"`
}
func (g SetItemInstallmentStatusResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
