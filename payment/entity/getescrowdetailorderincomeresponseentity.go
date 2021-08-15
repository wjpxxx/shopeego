package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetEscrowDetailOrderIncomeResponseEntity struct{
    EscrowAmount	float32	`json:"escrow_amount"`
    BuyerTotalAmount	float32	`json:"buyer_total_amount"`
    OriginalPrice	float32	`json:"original_price"`
    SellerDiscount	float32	`json:"seller_discount"`
    ShopeeDiscount	float32	`json:"shopee_discount"`
    VoucherFromSeller	float32	`json:"voucher_from_seller"`
    VoucherFromShopee	float32	`json:"voucher_from_shopee"`
    Coins	float32	`json:"coins"`
    BuyerPaidShippingFee	float32	`json:"buyer_paid_shipping_fee"`
    BuyerTransactionFee	float32	`json:"buyer_transaction_fee"`
    CrossBorderTax	float32	`json:"cross_border_tax"`
    PaymentPromotion	float32	`json:"payment_promotion"`
    CommissionFee	float32	`json:"commission_fee"`
    ServiceFee	float32	`json:"service_fee"`
    SellerTransactionFee	float32	`json:"seller_transaction_fee"`
    SellerLostCompensation	float32	`json:"seller_lost_compensation"`
    SellerCoinCashBack	float32	`json:"seller_coin_cash_back"`
    EscrowTax	float32	`json:"escrow_tax"`
    FinalShippingFee	float32	`json:"final_shipping_fee"`
    ActualShippingFee	float32	`json:"actual_shipping_fee"`
    ShopeeShippingRebate	float32	`json:"shopee_shipping_rebate"`
    ShippingFeeDiscountFrom3pl	float32	`json:"shipping_fee_discount_from_3pl"`
    SellerShippingDiscount	float32	`json:"seller_shipping_discount"`
    EstimatedShippingFee	float32	`json:"estimated_shipping_fee"`
    SellerVoucherCode	[]string	`json:"seller_voucher_code"`
    DrcAdjustableRefund	float32	`json:"drc_adjustable_refund"`
    CostOfGoodsSold	float32	`json:"cost_of_goods_sold"`
    OriginalCostOfGoodsSold	float32	`json:"original_cost_of_goods_sold"`
    OriginalShopeeDiscount	float32	`json:"original_shopee_discount"`
    SellerReturnRefund	float32	`json:"seller_return_refund"`
    ItemList	[]GetEscrowDetailItemListResponseEntity	`json:"item_list"`
    EscrowAmountPri	float32	`json:"escrow_amount_pri"`
    BuyerTotalAmountPri	float32	`json:"buyer_total_amount_pri"`
    OriginalPricePri	float32	`json:"original_price_pri"`
    SellerReturnRefundPri	float32	`json:"seller_return_refund_pri"`
    CommissionFeePri	float32	`json:"commission_fee_pri"`
    ServiceFeePri	float32	`json:"service_fee_pri"`
    DrcAdjustableRefundPri	float32	`json:"drc_adjustable_refund_pri"`
    PriCurrency	string	`json:"pri_currency"`
    AffCurrency	string	`json:"aff_currency"`
    ExchangeRate	float32	`json:"exchange_rate"`
}
func (g GetEscrowDetailOrderIncomeResponseEntity) String() string {
    return lib.ObjectToString(g)
}
