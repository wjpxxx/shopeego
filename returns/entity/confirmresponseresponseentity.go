package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ConfirmResponseResponseEntity struct{
    ReturnSn	string	`json:"return_sn"`
}
func (g ConfirmResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
