package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddTopPicksResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	AddTopPicksResponseResponseEntity	`json:"response"`
}
func (g AddTopPicksResult) String() string {
    return lib.ObjectToString(g)
}
