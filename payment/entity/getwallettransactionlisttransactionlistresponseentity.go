package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetWalletTransactionListTransactionListResponseEntity struct{
    Status	string	`json:"status"`
    TransactionType	string	`json:"transaction_type"`
    Amount	float32	`json:"amount"`
    CurrentBalance	float32	`json:"current_balance"`
    CreateTime	int	`json:"create_time"`
    OrderSn	string	`json:"order_sn"`
    RefundSn	string	`json:"refund_sn"`
    WithdrawalType	string	`json:"withdrawal_type"`
    TransactionFee	float32	`json:"transaction_fee"`
    Description	string	`json:"description"`
    BuyerName	string	`json:"buyer_name"`
    PayOrderList	[]GetWalletTransactionListPayOrderListResponseEntity	`json:"pay_order_list"`
    ShopName	string	`json:"shop_name"`
    WithdrawId	int64	`json:"withdraw_id"`
    Reason	string	`json:"reason"`
    RootWithdrawalId	int64	`json:"root_withdrawal_id"`
}
func (g GetWalletTransactionListTransactionListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
