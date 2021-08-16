package voucher

import (
	"github.com/wjpxxx/letgo/lib"
	shopeeConfig "github.com/wjpxxx/shopeego/config"
	voucherentity "github.com/wjpxxx/shopeego/voucher/entity"
)

//Voucher
type Voucher struct {
	Config *shopeeConfig.Config
}


//AddVoucher
//@Title Add voucher
//@Description https://open.shopee.com/documents?module=99&type=1&id=723&version=2
func (m *Voucher) AddVoucher(params voucherentity.AddVoucherRequest) voucherentity.AddVoucherResult {
    method := "voucher/add_voucher"
    result := voucherentity.AddVoucherResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//DeleteVoucher
//@Title Delete voucher
//@Description https://open.shopee.com/documents?module=99&type=1&id=724&version=2
func (m *Voucher) DeleteVoucher(params voucherentity.DeleteVoucherRequest) voucherentity.DeleteVoucherResult {
    method := "voucher/delete_voucher"
    result := voucherentity.DeleteVoucherResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//EndVoucher
//@Title End Voucher
//@Description https://open.shopee.com/documents?module=99&type=1&id=725&version=2
func (m *Voucher) EndVoucher(params voucherentity.EndVoucherRequest) voucherentity.EndVoucherResult {
    method := "voucher/end_voucher"
    result := voucherentity.EndVoucherResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//UpdateVoucher
//@Title Update voucher
//@Description https://open.shopee.com/documents?module=99&type=1&id=726&version=2
func (m *Voucher) UpdateVoucher(params voucherentity.UpdateVoucherRequest) voucherentity.UpdateVoucherResult {
    method := "voucher/update_voucher"
    result := voucherentity.UpdateVoucherResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//GetVoucherDetail
//@Title Get Voucher Detail
//@Description https://open.shopee.com/documents?module=99&type=1&id=727&version=2
func (m *Voucher) GetVoucherDetail(params voucherentity.GetVoucherDetailRequest) voucherentity.GetVoucherDetailResult {
    method := "voucher/get_voucher_detail"
    result := voucherentity.GetVoucherDetailResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//GetVoucherList
//@Title Get Voucher List
//@Description https://open.shopee.com/documents?module=99&type=1&id=728&version=2
func (m *Voucher) GetVoucherList(params voucherentity.GetVoucherListRequest) voucherentity.GetVoucherListResult {
    method := "voucher/get_voucher_list"
    result := voucherentity.GetVoucherListResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
