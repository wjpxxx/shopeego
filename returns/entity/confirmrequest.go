package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ConfirmRequest struct{
    ReturnSn	string	`json:"return_sn"`
}
func (g ConfirmRequest) String() string {
    return lib.ObjectToString(g)
}
