package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteAddOnDealSubItemSubItemListResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    ModelId	int64	`json:"model_id"`
    FailError	string	`json:"fail_error"`
    FailMessage	string	`json:"fail_message"`
}
func (g DeleteAddOnDealSubItemSubItemListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
