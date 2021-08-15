package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetItemInstallmentStatusItemInstallmentListResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    TenureList	[]int	`json:"tenure_list"`
}
func (g SetItemInstallmentStatusItemInstallmentListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
