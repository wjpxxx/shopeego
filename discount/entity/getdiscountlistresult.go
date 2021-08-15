package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetDiscountListResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	GetDiscountListResponseResponseEntity	`json:"response"`
    More	bool	`json:"more"`
}
func (g GetDiscountListResult) String() string {
    return lib.ObjectToString(g)
}
