package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReturnDetailRequest struct{
    ReturnSn	int	`json:"return_sn"`
}
func (g GetReturnDetailRequest) String() string {
    return lib.ObjectToString(g)
}
