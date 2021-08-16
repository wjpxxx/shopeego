package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ConfirmResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	ConfirmResponseResponseEntity	`json:"response"`
}
func (g ConfirmResult) String() string {
    return lib.ObjectToString(g)
}
