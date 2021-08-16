package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOfferToggleStatusRequest struct{
}
func (g GetOfferToggleStatusRequest) String() string {
    return lib.ObjectToString(g)
}
