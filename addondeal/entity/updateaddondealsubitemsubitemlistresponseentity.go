package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateAddOnDealSubItemSubItemListResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    Status	int	`json:"status"`
    ModelId	int64	`json:"model_id"`
    FailError	string	`json:"fail_error"`
    FailMessage	string	`json:"fail_message"`
    SubItemInputPrice	float32	`json:"sub_item_input_price"`
    SubItemLimit	int	`json:"sub_item_limit"`
}
func (g UpdateAddOnDealSubItemSubItemListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
