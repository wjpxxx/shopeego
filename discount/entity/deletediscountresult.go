package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteDiscountResult struct{
    DiscountId	int64	`json:"discount_id"`
    ModifyTime	int	`json:"modify_time"`
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Msg	string	`json:"msg"`
}
func (g DeleteDiscountResult) String() string {
    return lib.ObjectToString(g)
}
