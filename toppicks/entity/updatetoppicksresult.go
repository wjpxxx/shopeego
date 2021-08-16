package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateTopPicksResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	UpdateTopPicksResponseResponseEntity	`json:"response"`
    RequestId	string	`json:"request_id"`
}
func (g UpdateTopPicksResult) String() string {
    return lib.ObjectToString(g)
}
