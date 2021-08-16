package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateAddOnDealSubItemResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	UpdateAddOnDealSubItemResponseResponseEntity	`json:"response"`
    AddOnDealId	int64	`json:"add_on_deal_id"`
}
func (g UpdateAddOnDealSubItemResult) String() string {
    return lib.ObjectToString(g)
}
