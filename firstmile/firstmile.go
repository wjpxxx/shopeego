package firstmile

import (
	"github.com/wjpxxx/letgo/lib"
	shopeeConfig "github.com/wjpxxx/shopeego/config"
	firstmileentity "github.com/wjpxxx/shopeego/firstmile/entity"
)

//FirstMile
type FirstMile struct {
	Config *shopeeConfig.Config
}

//GetUnbindOrderList
//@Title Use this api to get unbind order list.
//@Description https://open.shopee.com/documents?module=96&type=1&id=605&version=2
func (m *FirstMile) GetUnbindOrderList(params firstmileentity.GetUnbindOrderListRequest) firstmileentity.GetUnbindOrderListResult {
    method := "first_mile/get_unbind_order_list"
    result := firstmileentity.GetUnbindOrderListResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//GetDetail
//@Title Use this api to get first mile detail.
//@Description https://open.shopee.com/documents?module=96&type=1&id=601&version=2
func (m *FirstMile) GetDetail(params firstmileentity.GetDetailRequest) firstmileentity.GetDetailResult {
    method := "first_mile/get_detail"
    result := firstmileentity.GetDetailResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//GenerateFirstMileTrackingNumber
//@Title Use this api to generate first mile tracking number.
//@Description https://open.shopee.com/documents?module=96&type=1&id=600&version=2
func (m *FirstMile) GenerateFirstMileTrackingNumber(params firstmileentity.GenerateFirstMileTrackingNumberRequest) firstmileentity.GenerateFirstMileTrackingNumberResult {
    method := "first_mile/generate_first_mile_tracking_number"
    result := firstmileentity.GenerateFirstMileTrackingNumberResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//BindFirstMileTrackingNumber
//@Title Use this api to bind first mile tracking number.
//@Description https://open.shopee.com/documents?module=96&type=1&id=599&version=2
func (m *FirstMile) BindFirstMileTrackingNumber(params firstmileentity.BindFirstMileTrackingNumberRequest) firstmileentity.BindFirstMileTrackingNumberResult {    
	method := "first_mile/bind_first_mile_tracking_number"
    result := firstmileentity.BindFirstMileTrackingNumberResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//UnbindFirstMileTrackingNumber
//@Title Use this api to unbind first mile.
//@Description https://open.shopee.com/documents?module=96&type=1&id=604&version=2
func (m *FirstMile) UnbindFirstMileTrackingNumber(params firstmileentity.UnbindFirstMileTrackingNumberRequest) firstmileentity.UnbindFirstMileTrackingNumberResult {
    method := "first_mile/unbind_first_mile_tracking_number"
    result := firstmileentity.UnbindFirstMileTrackingNumberResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//GetTrackingNumberList
//@Title Use this api to get first mile tracking number list.
//@Description https://open.shopee.com/documents?module=96&type=1&id=602&version=2
func (m *FirstMile) GetTrackingNumberList(params firstmileentity.GetTrackingNumberListRequest) firstmileentity.GetTrackingNumberListResult {
    method := "first_mile/get_tracking_number_list"
    result := firstmileentity.GetTrackingNumberListResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//GetWaybill
//@Title Use this api to get first mile waybill file.
//@Description https://open.shopee.com/documents?module=96&type=1&id=603&version=2
func (m *FirstMile) GetWaybill(params firstmileentity.GetWaybillRequest) firstmileentity.GetWaybillResult {
    method := "first_mile/get_waybill"
    result := firstmileentity.GetWaybillResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//FGetChannelList
//@Title Use this api to get first mile channel list.
//@Description https://open.shopee.com/documents?module=96&type=1&id=606&version=2
func (m *FirstMile) FGetChannelList(params firstmileentity.GetChannelListRequest) firstmileentity.GetChannelListResult {
    method := "first_mile/get_channel_list"
    result := firstmileentity.GetChannelListResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}