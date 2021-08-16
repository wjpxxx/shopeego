package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetPushConfigPushConfigRequestEntity struct{
    OrderStatus	int	`json:"order_status"`
    OrderTrackingNo	int	`json:"order_tracking_no"`
    ShopUpdate	int	`json:"shop_update"`
    BannedItem	int	`json:"banned_item"`
    ItmePromotion	int	`json:"itme_promotion"`
    ReservedStockChange	int	`json:"reserved_stock_change"`
    Webchat	int	`json:"webchat"`
}
func (g SetPushConfigPushConfigRequestEntity) String() string {
    return lib.ObjectToString(g)
}
