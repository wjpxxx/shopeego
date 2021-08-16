package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddAddOnDealSubItemSubItemListResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    Status	int	`json:"status"`
    ModelId	int64	`json:"model_id"`
    FailError	string	`json:"fail_error"`
    FailMessage	string	`json:"fail_message"`
}
func (g AddAddOnDealSubItemSubItemListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
