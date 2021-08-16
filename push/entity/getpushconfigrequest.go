package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetPushConfigRequest struct{
}
func (g GetPushConfigRequest) String() string {
    return lib.ObjectToString(g)
}
