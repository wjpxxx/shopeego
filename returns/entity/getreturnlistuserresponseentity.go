package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReturnListUserResponseEntity struct{
    Username	string	`json:"username"`
    Email	string	`json:"email"`
    Portrait	string	`json:"portrait"`
}
func (g GetReturnListUserResponseEntity) String() string {
    return lib.ObjectToString(g)
}
