package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetPaymentMethodListResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	[]GetPaymentMethodListResponseResponseEntity	`json:"response"`
}
func (g GetPaymentMethodListResult) String() string {
    return lib.ObjectToString(g)
}
