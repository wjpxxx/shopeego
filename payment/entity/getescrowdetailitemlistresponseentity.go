package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetEscrowDetailItemListResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    ItemName	string	`json:"item_name"`
    ItemSku	string	`json:"item_sku"`
    ModelId	int64	`json:"model_id"`
    ModelName	string	`json:"model_name"`
    ModelSku	string	`json:"model_sku"`
    OriginalPrice	float32	`json:"original_price"`
    DiscountedPrice	float32	`json:"discounted_price"`
    DiscountFromCoin	float32	`json:"discount_from_coin"`
    DiscountFromVoucherShopee	float32	`json:"discount_from_voucher_shopee"`
    DiscountFromVoucherSeller	float32	`json:"discount_from_voucher_seller"`
    ActivityType	string	`json:"activity_type"`
    ActivityId	int64	`json:"activity_id"`
    IsMainItem	bool	`json:"is_main_item"`
}
func (g GetEscrowDetailItemListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
