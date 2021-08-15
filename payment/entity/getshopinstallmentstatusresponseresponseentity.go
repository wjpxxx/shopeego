package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetShopInstallmentStatusResponseResponseEntity struct{
    InstallmentStatus	int	`json:"installment_status"`
}
func (g GetShopInstallmentStatusResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
