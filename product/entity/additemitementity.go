package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//AddItemItemEntity
type AddItemItemEntity struct{
	ItemID int64 `json:"item_id"`
	CategoryID int64 `json:"category_id"`
	ItemName string `json:"item_name"`
	Description string `json:"description"`
	ItemSku string `json:"item_sku"`
	CreateTime int `json:"create_time"`
	UpdateTime int `json:"update_time"`
	Attributes []AttributeEntity `json:"attributes"`
	PriceInfo PriceInfoEntity `json:"price_info"`
	StockInfo StockInfoEntity `json:"stock_info"`
	Image ImageEntity `json:"image"`
	Images ImageEntity  `json:"images"`
	Weight string `json:"weight"`
	Dimension DimensionEntity `json:"dimension"`
	LogisticInfo []LogisticInfoEntity `json:"logistic_info"`
	PreOrder PreOrderEntity `json:"pre_order"`
	Wholesales []WholesalesEntity `json:"wholesales"`
	Condition string `json:"condition"`
	SizeChart string `json:"size_chart"`
	ItemStatus string `json:"item_status"`
	HasModel bool `json:"has_model"`
	PromotionID int64 `json:"promotion_id"`
	VideoInfo []VideoInfoEntity `json:"video_info"`
	Brand BrandEntity `json:"brand"`
	ItemDangerous int `json:"item_dangerous"`
}

//String
func(i AddItemItemEntity)String()string{
	return lib.ObjectToString(i)
}