package firstmile

import (
	"github.com/wjpxxx/letgo/lib"
	shopeeConfig "github.com/wjpxxx/shopeego/config"
	firstmileEntity "github.com/wjpxxx/shopeego/firstmile/entity"
)

//FirstMile
type FirstMile struct {
	Config *shopeeConfig.Config
}

//GetUnbindOrderList
//@Title Use this api to get unbind order list.
//@Description https://open.shopee.com/documents?module=96&type=1&id=605&version=2
func (m *FirstMile) GetUnbindOrderList(cursor string, pageSize int, responseOptionalFields string) firstmileEntity.GetUnbindOrderListResult {
	method := "first_mile/get_unbind_order_list"
	result := firstmileEntity.GetUnbindOrderListResult{}
	params := lib.InRow{
		"cursor":                   cursor,
		"page_size":                pageSize,
		"response_optional_fields": responseOptionalFields,
	}
	err := m.Config.HttpGet(method, params, &result)
	if err != nil {
		result.Error = err.Error()
	}
	return result
}
