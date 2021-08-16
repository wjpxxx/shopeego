package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteTopPicksResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	DeleteTopPicksResponseResponseEntity	`json:"response"`
}
func (g DeleteTopPicksResult) String() string {
    return lib.ObjectToString(g)
}
