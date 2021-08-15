package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetShopInstallmentStatusResponseResponseEntity struct{
    InstallmentStatus	int	`json:"installment_status"`
}
func (g SetShopInstallmentStatusResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
