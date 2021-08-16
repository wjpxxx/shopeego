package bundledeal

import (
	"github.com/wjpxxx/letgo/lib"
	shopeeConfig "github.com/wjpxxx/shopeego/config"
	bundledealentity "github.com/wjpxxx/shopeego/bundledeal/entity"
)

//BundleDeal
type BundleDeal struct {
	Config *shopeeConfig.Config
}


//AddBundleDeal
//@Title create bundle deal
//@Description https://open.shopee.com/documents?module=99&type=1&id=689&version=2
func (m *BundleDeal) AddBundleDeal(params bundledealentity.AddBundleDealRequest) bundledealentity.AddBundleDealResult {
    method := "bundle_deal/add_bundle_deal"
    result := bundledealentity.AddBundleDealResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//AddBundleDealItem
//@Title add product to bundle deal
//@Description https://open.shopee.com/documents?module=99&type=1&id=690&version=2
func (m *BundleDeal) AddBundleDealItem(params bundledealentity.AddBundleDealItemRequest) bundledealentity.AddBundleDealItemResult {
    method := "bundle_deal/add_bundle_deal_item"
    result := bundledealentity.AddBundleDealItemResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//DeleteBundleDeal
//@Title delete bundle deal
//@Description https://open.shopee.com/documents?module=99&type=1&id=691&version=2
func (m *BundleDeal) DeleteBundleDeal(params bundledealentity.DeleteBundleDealRequest) bundledealentity.DeleteBundleDealResult {
    method := "bundle_deal/delete_bundle_deal"
    result := bundledealentity.DeleteBundleDealResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//DeleteBundleDealItem
//@Title delete product in bundle deal
//@Description https://open.shopee.com/documents?module=99&type=1&id=692&version=2
func (m *BundleDeal) DeleteBundleDealItem(params bundledealentity.DeleteBundleDealItemRequest) bundledealentity.DeleteBundleDealItemResult {
    method := "bundle_deal/delete_bundle_deal_item"
    result := bundledealentity.DeleteBundleDealItemResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//EndBundleDeal
//@Title end bundle deal
//@Description https://open.shopee.com/documents?module=99&type=1&id=693&version=2
func (m *BundleDeal) EndBundleDeal(params bundledealentity.EndBundleDealRequest) bundledealentity.EndBundleDealResult {
    method := "bundle_deal/end_bundle_deal"
    result := bundledealentity.EndBundleDealResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//GetBundleDealList
//@Title get bundle deal list
//@Description https://open.shopee.com/documents?module=99&type=1&id=694&version=2
func (m *BundleDeal) GetBundleDealList(params bundledealentity.GetBundleDealListRequest) bundledealentity.GetBundleDealListResult {
    method := "bundle_deal/get_bundle_deal_list"
    result := bundledealentity.GetBundleDealListResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//GetBundleDeal
//@Title get bundle deal detail
//@Description https://open.shopee.com/documents?module=99&type=1&id=695&version=2
func (m *BundleDeal) GetBundleDeal(params bundledealentity.GetBundleDealRequest) bundledealentity.GetBundleDealResult {
    method := "bundle_deal/get_bundle_deal"
    result := bundledealentity.GetBundleDealResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//GetBundleDealItem
//@Title get bundle deal item
//@Description https://open.shopee.com/documents?module=99&type=1&id=696&version=2
func (m *BundleDeal) GetBundleDealItem(params bundledealentity.GetBundleDealItemRequest) bundledealentity.GetBundleDealItemResult {
    method := "bundle_deal/get_bundle_deal_item"
    result := bundledealentity.GetBundleDealItemResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//UpdateBundleDeal
//@Title update bundle deal
//@Description https://open.shopee.com/documents?module=99&type=1&id=697&version=2
func (m *BundleDeal) UpdateBundleDeal(params bundledealentity.UpdateBundleDealRequest) bundledealentity.UpdateBundleDealResult {
    method := "bundle_deal/update_bundle_deal"
    result := bundledealentity.UpdateBundleDealResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//UpdateBundleDealItem
//@Title update product in bundle deal
//@Description https://open.shopee.com/documents?module=99&type=1&id=698&version=2
func (m *BundleDeal) UpdateBundleDealItem(params bundledealentity.UpdateBundleDealItemRequest) bundledealentity.UpdateBundleDealItemResult {
    method := "bundle_deal/update_bundle_deal_item"
    result := bundledealentity.UpdateBundleDealItemResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}