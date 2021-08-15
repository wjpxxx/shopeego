package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetPaymentMethodListRequest struct{
}
func (g GetPaymentMethodListRequest) String() string {
    return lib.ObjectToString(g)
}
