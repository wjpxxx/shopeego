package addondeal

import (
	"github.com/wjpxxx/letgo/lib"
	shopeeConfig "github.com/wjpxxx/shopeego/config"
	addondealentity "github.com/wjpxxx/shopeego/addondeal/entity"
)

//AddOnDeal
type AddOnDeal struct {
	Config *shopeeConfig.Config
}

//AddAddOnDealMainItem
//@Title Add Add-on Deal Main Item
//@Description https://open.shopee.com/documents?module=99&type=1&id=708&version=2
func (m *AddOnDeal) AddAddOnDealMainItem(params addondealentity.AddAddOnDealMainItemRequest) addondealentity.AddAddOnDealMainItemResult {
    method := "add_on_deal/add_add_on_deal_main_item"
    result := addondealentity.AddAddOnDealMainItemResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//AddAddOnDealSubItem
//@Title Add Add-on Deal Sub Item
//@Description https://open.shopee.com/documents?module=99&type=1&id=709&version=2
func (m *AddOnDeal) AddAddOnDealSubItem(params addondealentity.AddAddOnDealSubItemRequest) addondealentity.AddAddOnDealSubItemResult {
    method := "add_on_deal/add_add_on_deal_sub_item"
    result := addondealentity.AddAddOnDealSubItemResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//AddAddOnDeal
//@Title Add Add-on Deal
//@Description https://open.shopee.com/documents?module=99&type=1&id=710&version=2
func (m *AddOnDeal) AddAddOnDeal(params addondealentity.AddAddOnDealRequest) addondealentity.AddAddOnDealResult {
    method := "add_on_deal/add_add_on_deal"
    result := addondealentity.AddAddOnDealResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//DeleteAddOnDealMainItem
//@Title Delete Add-on Deal Main Item
//@Description https://open.shopee.com/documents?module=99&type=1&id=711&version=2
func (m *AddOnDeal) DeleteAddOnDealMainItem(params addondealentity.DeleteAddOnDealMainItemRequest) addondealentity.DeleteAddOnDealMainItemResult {
    method := "add_on_deal/delete_add_on_deal_main_item"
    result := addondealentity.DeleteAddOnDealMainItemResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//DeleteAddOnDealSubItem
//@Title Delete Add-on Deal Sub Item
//@Description https://open.shopee.com/documents?module=99&type=1&id=712&version=2
func (m *AddOnDeal) DeleteAddOnDealSubItem(params addondealentity.DeleteAddOnDealSubItemRequest) addondealentity.DeleteAddOnDealSubItemResult {
    method := "add_on_deal/delete_add_on_deal_sub_item"
    result := addondealentity.DeleteAddOnDealSubItemResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//DeleteAddOnDeal
//@Title Delete Add-on Deal
//@Description https://open.shopee.com/documents?module=99&type=1&id=713&version=2
func (m *AddOnDeal) DeleteAddOnDeal(params addondealentity.DeleteAddOnDealRequest) addondealentity.DeleteAddOnDealResult {
    method := "add_on_deal/delete_add_on_deal"
    result := addondealentity.DeleteAddOnDealResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//EndAddOnDeal
//@Title End Add-on Deal
//@Description https://open.shopee.com/documents?module=99&type=1&id=714&version=2
func (m *AddOnDeal) EndAddOnDeal(params addondealentity.EndAddOnDealRequest) addondealentity.EndAddOnDealResult {
    method := "add_on_deal/end_add_on_deal"
    result := addondealentity.EndAddOnDealResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//GetAddOnDealList
//@Title Get Add-on Deal List
//@Description https://open.shopee.com/documents?module=99&type=1&id=715&version=2
func (m *AddOnDeal) GetAddOnDealList(params addondealentity.GetAddOnDealListRequest) addondealentity.GetAddOnDealListResult {
    method := "add_on_deal/get_add_on_deal_list"
    result := addondealentity.GetAddOnDealListResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//GetAddOnDealMainItem
//@Title Get Add-on Deal Main Item
//@Description https://open.shopee.com/documents?module=99&type=1&id=717&version=2
func (m *AddOnDeal) GetAddOnDealMainItem(params addondealentity.GetAddOnDealMainItemRequest) addondealentity.GetAddOnDealMainItemResult {
    method := "add_on_deal/get_add_on_deal_main_item"
    result := addondealentity.GetAddOnDealMainItemResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//GetAddOnDealSubItem
//@Title Get Add-on Deal Sub Item
//@Description https://open.shopee.com/documents?module=99&type=1&id=718&version=2
func (m *AddOnDeal) GetAddOnDealSubItem(params addondealentity.GetAddOnDealSubItemRequest) addondealentity.GetAddOnDealSubItemResult {
    method := "add_on_deal/get_add_on_deal_sub_item"
    result := addondealentity.GetAddOnDealSubItemResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//GetAddOnDeal
//@Title Get Add-on Deal
//@Description https://open.shopee.com/documents?module=99&type=1&id=719&version=2
func (m *AddOnDeal) GetAddOnDeal(params addondealentity.GetAddOnDealRequest) addondealentity.GetAddOnDealResult {
    method := "add_on_deal/get_add_on_deal"
    result := addondealentity.GetAddOnDealResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//UpdateAddOnDealMainItem
//@Title Update Add-on Deal Main Item
//@Description https://open.shopee.com/documents?module=99&type=1&id=720&version=2
func (m *AddOnDeal) UpdateAddOnDealMainItem(params addondealentity.UpdateAddOnDealMainItemRequest) addondealentity.UpdateAddOnDealMainItemResult {
    method := "add_on_deal/update_add_on_deal_main_item"
    result := addondealentity.UpdateAddOnDealMainItemResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//UpdateAddOnDealSubItem
//@Title Update Add-on Deal Sub Item
//@Description https://open.shopee.com/documents?module=99&type=1&id=721&version=2
func (m *AddOnDeal) UpdateAddOnDealSubItem(params addondealentity.UpdateAddOnDealSubItemRequest) addondealentity.UpdateAddOnDealSubItemResult {
    method := "add_on_deal/update_add_on_deal_sub_item"
    result := addondealentity.UpdateAddOnDealSubItemResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//UpdateAddOnDeal
//@Title Update Add-on Deal
//@Description https://open.shopee.com/documents?module=99&type=1&id=722&version=2
func (m *AddOnDeal) UpdateAddOnDeal(params addondealentity.UpdateAddOnDealRequest) addondealentity.UpdateAddOnDealResult {
    method := "add_on_deal/update_add_on_deal"
    result := addondealentity.UpdateAddOnDealResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

