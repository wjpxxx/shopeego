package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReturnDetailItemResponseEntity struct{
    ModelId	int64	`json:"model_id"`
    Name	string	`json:"name"`
    Images	[]string	`json:"images"`
    Amount	int	`json:"amount"`
    ItemPrice	float32	`json:"item_price"`
    IsAddOnDeal	bool	`json:"is_add_on_deal"`
    IsMainItem	bool	`json:"is_main_item"`
    AddOnDealId	int64	`json:"add_on_deal_id"`
    ItemId	int64	`json:"item_id"`
    ItemSku	string	`json:"item_sku"`
    VariationSku	string	`json:"variation_sku"`
}
func (g GetReturnDetailItemResponseEntity) String() string {
    return lib.ObjectToString(g)
}
