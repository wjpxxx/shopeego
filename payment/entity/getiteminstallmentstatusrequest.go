package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetItemInstallmentStatusRequest struct{
    ItemIdList	[]int64	`json:"item_id_list"`
}
func (g GetItemInstallmentStatusRequest) String() string {
    return lib.ObjectToString(g)
}
