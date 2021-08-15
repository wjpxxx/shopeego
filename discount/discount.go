package discount

import (
	"github.com/wjpxxx/letgo/lib"
	shopeeConfig "github.com/wjpxxx/shopeego/config"
	discountentity "github.com/wjpxxx/shopeego/discount/entity"
)

//Discount
type Discount struct {
	Config *shopeeConfig.Config
}

//AddDiscount
//@Title Use this api to add shop discount activity
//@Description https://open.shopee.com/documents?module=99&type=1&id=569&version=2
func (m *Discount) AddDiscount(params discountentity.AddDiscountRequest) discountentity.AddDiscountResult {
    method := "discount/add_discount"
    result := discountentity.AddDiscountResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//AddDiscountItem
//@Title Use this api to add shop discount item
//@Description https://open.shopee.com/documents?module=99&type=1&id=570&version=2
func (m *Discount) AddDiscountItem(params discountentity.AddDiscountItemRequest) discountentity.AddDiscountItemResult {
    method := "discount/add_discount_item"
    result := discountentity.AddDiscountItemResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//DeleteDiscount
//@Title Use this api to delete one discount activity
//@Description https://open.shopee.com/documents?module=99&type=1&id=571&version=2
func (m *Discount) DeleteDiscount(params discountentity.DeleteDiscountRequest) discountentity.DeleteDiscountResult {
    method := "discount/delete_discount"
    result := discountentity.DeleteDiscountResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//DeleteDiscountItem
//@Title Use this api to delete items of the discount activity
//@Description https://open.shopee.com/documents?module=99&type=1&id=572&version=2
func (m *Discount) DeleteDiscountItem(params discountentity.DeleteDiscountItemRequest) discountentity.DeleteDiscountItemResult {
    method := "discount/delete_discount_item"
    result := discountentity.DeleteDiscountItemResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//GetDiscount
//@Title Use this api to get one shop discount activity detail
//@Description https://open.shopee.com/documents?module=99&type=1&id=574&version=2
func (m *Discount) GetDiscount(params discountentity.GetDiscountRequest) discountentity.GetDiscountResult {
    method := "discount/get_discount"
    result := discountentity.GetDiscountResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//GetDiscountList
//@Title Use this api to get shop discount activity list
//@Description https://open.shopee.com/documents?module=99&type=1&id=575&version=2
func (m *Discount) GetDiscountList(params discountentity.GetDiscountListRequest) discountentity.GetDiscountListResult {
    method := "discount/get_discount_list"
    result := discountentity.GetDiscountListResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//UpdateDiscount
//@Title Use this api to update one discount information
//@Description https://open.shopee.com/documents?module=99&type=1&id=576&version=2
func (m *Discount) UpdateDiscount(params discountentity.UpdateDiscountRequest) discountentity.UpdateDiscountResult {
    method := "discount/update_discount"
    result := discountentity.UpdateDiscountResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//UpdateDiscountItem
//@Title Use this api to update items of the discount activity
//@Description https://open.shopee.com/documents?module=99&type=1&id=577&version=2
func (m *Discount) UpdateDiscountItem(params discountentity.UpdateDiscountItemRequest) discountentity.UpdateDiscountItemResult {
    method := "discount/update_discount_item"
    result := discountentity.UpdateDiscountItemResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

//EndDiscount
//@Title Use this api to end shop discount activity
//@Description https://open.shopee.com/documents?module=99&type=1&id=597&version=2
func (m *Discount) EndDiscount(params discountentity.EndDiscountRequest) discountentity.EndDiscountResult {
    method := "discount/end_discount"
    result := discountentity.EndDiscountResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}