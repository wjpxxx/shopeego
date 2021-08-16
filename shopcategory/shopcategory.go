package shopcategory

import (
	"github.com/wjpxxx/letgo/lib"
	shopeeConfig "github.com/wjpxxx/shopeego/config"
	shopcategoryentity "github.com/wjpxxx/shopeego/shopcategory/entity"
)

//ShopCategory
type ShopCategory struct {
	Config *shopeeConfig.Config
}

//AddShopCategory
//@Title Use this call to add a new shop collecion
//@Description https://open.shopee.com/documents?module=99&type=1&id=586&version=2
func (m *ShopCategory) AddShopCategory(params shopcategoryentity.AddShopCategoryRequest) shopcategoryentity.AddShopCategoryResult {
    method := "shop_category/add_shop_category"
    result := shopcategoryentity.AddShopCategoryResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//GetShopCategoryList
//@Title Use this call to get list of shop categories
//@Description https://open.shopee.com/documents?module=99&type=1&id=587&version=2
func (m *ShopCategory) GetShopCategoryList(params shopcategoryentity.GetShopCategoryListRequest) shopcategoryentity.GetShopCategoryListResult {
    method := "shop_category/get_shop_category_list"
    result := shopcategoryentity.GetShopCategoryListResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//DeleteShopCategory
//@Title Use this call to delete a existing shop collecion
//@Description https://open.shopee.com/documents?module=99&type=1&id=588&version=2
func (m *ShopCategory) DeleteShopCategory(params shopcategoryentity.DeleteShopCategoryRequest) shopcategoryentity.DeleteShopCategoryResult {
    method := "shop_category/delete_shop_category"
    result := shopcategoryentity.DeleteShopCategoryResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//UpdateShopCategory
//@Title Use this call to update a existing collecion
//@Description https://open.shopee.com/documents?module=99&type=1&id=589&version=2
func (m *ShopCategory) UpdateShopCategory(params shopcategoryentity.UpdateShopCategoryRequest) shopcategoryentity.UpdateShopCategoryResult {
    method := "shop_category/update_shop_category"
    result := shopcategoryentity.UpdateShopCategoryResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//AddItemList
//@Title Use this call to add items list to certain shop_category
//@Description https://open.shopee.com/documents?module=99&type=1&id=590&version=2
func (m *ShopCategory) AddItemList(params shopcategoryentity.AddItemListRequest) shopcategoryentity.AddItemListResult {
    method := "shop_category/add_item_list"
    result := shopcategoryentity.AddItemListResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//CGetItemList
//@Title Use this call to get items list of certain shop_category
//@Description https://open.shopee.com/documents?module=99&type=1&id=591&version=2
func (m *ShopCategory) CGetItemList(params shopcategoryentity.GetItemListRequest) shopcategoryentity.GetItemListResult {
    method := "shop_category/get_item_list"
    result := shopcategoryentity.GetItemListResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//DeleteItemList
//@Title Use this api to delete items from shop category
//@Description https://open.shopee.com/documents?module=99&type=1&id=592&version=2
func (m *ShopCategory) DeleteItemList(params shopcategoryentity.DeleteItemListRequest) shopcategoryentity.DeleteItemListResult {
    method := "shop_category/delete_item_list"
    result := shopcategoryentity.DeleteItemListResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
