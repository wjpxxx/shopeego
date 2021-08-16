package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetAddOnDealSubItemSubItemListResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    Status	int	`json:"status"`
    SubItemLimit	int	`json:"sub_item_limit"`
    ModelId	int64	`json:"model_id"`
    Price	GetAddOnDealSubItemPriceResponseEntity	`json:"price"`
}
func (g GetAddOnDealSubItemSubItemListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
