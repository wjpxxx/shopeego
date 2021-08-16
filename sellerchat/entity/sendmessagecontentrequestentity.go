package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SendMessageContentRequestEntity struct{
    Text	string	`json:"text"`
    StickerId	string	`json:"sticker_id"`
    StickerPackageId	string	`json:"sticker_package_id"`
    ImageUrl	string	`json:"image_url"`
    ItemId	int64	`json:"item_id"`
    OrderSn	string	`json:"order_sn"`
}
func (g SendMessageContentRequestEntity) String() string {
    return lib.ObjectToString(g)
}
