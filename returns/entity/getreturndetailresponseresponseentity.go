package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReturnDetailResponseResponseEntity struct{
    Image	[]string	`json:"image"`
    Reason	string	`json:"reason"`
    TextReason	string	`json:"text_reason"`
    ReturnSn	int	`json:"return_sn"`
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
    User	GetReturnDetailUserResponseEntity	`json:"user"`
    Item	[]GetReturnDetailItemResponseEntity	`json:"item"`
    OrderSn	string	`json:"order_sn"`
    ReturnShipDueDate	int	`json:"return_ship_due_date"`
    ReturnSellerDueDate	int	`json:"return_seller_due_date"`
    Activity	[]GetReturnDetailActivityResponseEntity	`json:"activity"`
}
func (g GetReturnDetailResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
