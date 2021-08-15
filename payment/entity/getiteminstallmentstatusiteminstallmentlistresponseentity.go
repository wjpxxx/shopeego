package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetItemInstallmentStatusItemInstallmentListResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    TenureList	[]int	`json:"tenure_list"`
}
func (g GetItemInstallmentStatusItemInstallmentListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
