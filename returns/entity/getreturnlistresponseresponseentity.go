package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReturnListResponseResponseEntity struct{
    Image	[]string	`json:"image"`
    Reason	string	`json:"reason"`
    TextReason	string	`json:"text_reason"`
    ReturnSn	int64	`json:"return_sn"`
    RefundAmount	float32	`json:"refund_amount"`
    Currency	string	`json:"currency"`
    CreateTime	int	`json:"create_time"`
    UpdateTime	int	`json:"update_time"`
    Status	string	`json:"status"`
    DueDate	int	`json:"due_date"`
    TrackingNumber	string	`json:"tracking_number"`
    DisputeReason	[]string	`json:"dispute_reason"`
    DisputeTextReason	[]string	`json:"dispute_text_reason"`
    NeedsLogistics	bool	`json:"needs_logistics"`
    AmountBeforeDiscount	float32	`json:"amount_before_discount"`
    User	GetReturnListUserResponseEntity	`json:"user"`
    Item	[]GetReturnListItemResponseEntity	`json:"item"`
    OrderSn	string	`json:"order_sn"`
    ReturnShipDueDate	int	`json:"return_ship_due_date"`
    ReturnSellerDueDate	int	`json:"return_seller_due_date"`
}
func (g GetReturnListResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
