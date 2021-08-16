package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReturnDetailUserResponseEntity struct{
    Username	string	`json:"username"`
    Email	string	`json:"email"`
    Portrait	string	`json:"portrait"`
}
func (g GetReturnDetailUserResponseEntity) String() string {
    return lib.ObjectToString(g)
}
