package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetShopInstallmentStatusRequest struct{
}
func (g GetShopInstallmentStatusRequest) String() string {
    return lib.ObjectToString(g)
}
