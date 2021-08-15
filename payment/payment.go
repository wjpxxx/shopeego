package payment

import (
	"github.com/wjpxxx/letgo/lib"
	shopeeConfig "github.com/wjpxxx/shopeego/config"
	paymententity "github.com/wjpxxx/shopeego/payment/entity"
)

//Payment
type Payment struct {
	Config *shopeeConfig.Config
}

//GetEscrowDetail
//@Title Use this API to fetch the accounting detail of order.
//@Description https://open.shopee.com/documents?module=97&type=1&id=565&version=2
func (m *Payment) GetEscrowDetail(params paymententity.GetEscrowDetailRequest) paymententity.GetEscrowDetailResult {
    method := "payment/get_escrow_detail"
    result := paymententity.GetEscrowDetailResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//SetShopInstallmentStatus
//@Title Sets the staging capability of shop level.
//@Description https://open.shopee.com/documents?module=97&type=1&id=566&version=2
func (m *Payment) SetShopInstallmentStatus(params paymententity.SetShopInstallmentStatusRequest) paymententity.SetShopInstallmentStatusResult {
    method := "payment/set_shop_installment_status"
    result := paymententity.SetShopInstallmentStatusResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//GetShopInstallmentStatus
//@Title Get the installment state of shop.
//@Description https://open.shopee.com/documents?module=97&type=1&id=567&version=2
func (m *Payment) GetShopInstallmentStatus() paymententity.GetShopInstallmentStatusResult {
    method := "payment/get_shop_installment_status"
    result := paymententity.GetShopInstallmentStatusResult{}
	params := lib.InRow{}
    err := m.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//GetPayoutDetail
//@Title Get the payoutdetail for CB.
//@Description https://open.shopee.com/documents?module=97&type=1&id=573&version=2
func (m *Payment) GetPayoutDetail(params paymententity.GetPayoutDetailRequest) paymententity.GetPayoutDetailResult {
    method := "payment/get_payout_detail"
    result := paymententity.GetPayoutDetailResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//SetItemInstallmentStatus
//@Title Set item installment.Only for TH、TW.
//@Description https://open.shopee.com/documents?module=97&type=1&id=582&version=2
func (m *Payment) SetItemInstallmentStatus(params paymententity.SetItemInstallmentStatusRequest) paymententity.SetItemInstallmentStatusResult {
    method := "payment/set_item_installment_status"
    result := paymententity.SetItemInstallmentStatusResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//GetItemInstallmentStatus
//@Title Get item installment tenures.Only for TH、TW.
//@Description https://open.shopee.com/documents?module=97&type=1&id=583&version=2
func (m *Payment) GetItemInstallmentStatus(params paymententity.GetItemInstallmentStatusRequest) paymententity.GetItemInstallmentStatusResult {
    method := "payment/get_item_installment_status"
    result := paymententity.GetItemInstallmentStatusResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//GetPaymentMethodList
//@Title Obtain payment method (no authentication required)
//@Description https://open.shopee.com/documents?module=97&type=1&id=593&version=2
func (m *Payment) GetPaymentMethodList(params paymententity.GetPaymentMethodListRequest) paymententity.GetPaymentMethodListResult {
    method := "payment/get_payment_method_list"
    result := paymententity.GetPaymentMethodListResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//GetWalletTransactionList
//@Title Use this API to get the transaction records of wallet.
//@Description https://open.shopee.com/documents?module=97&type=1&id=594&version=2
func (m *Payment) GetWalletTransactionList(params paymententity.GetWalletTransactionListRequest) paymententity.GetWalletTransactionListResult {
    method := "payment/get_wallet_transaction_list"
    result := paymententity.GetWalletTransactionListResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//GetEscrowList
//@Title Use this API to fetch the accounting list of order.
//@Description https://open.shopee.com/documents?module=97&type=1&id=669&version=2
func (m *Payment) GetEscrowList(params paymententity.GetEscrowListRequest) paymententity.GetEscrowListResult {
    method := "payment/get_escrow_list"
    result := paymententity.GetEscrowListResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}