package accounthealth

import (
	"github.com/wjpxxx/letgo/lib"
	shopeeConfig "github.com/wjpxxx/shopeego/config"
	accounthealthentity "github.com/wjpxxx/shopeego/accounthealth/entity"
)

//AccountHealth
type AccountHealth struct {
	Config *shopeeConfig.Config
}

//ShopPerformance
//@Title The data metrics of shop performance
//@Description https://open.shopee.com/documents?module=99&type=1&id=658&version=2
func (m *AccountHealth) ShopPerformance() accounthealthentity.ShopPerformanceResult {
    method := "account_health/shop_performance"
	params := lib.InRow{}
    result := accounthealthentity.ShopPerformanceResult{}
    err := m.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//ShopPenalty
//@Title To get the information of shop penalty.
//@Description https://open.shopee.com/documents?module=99&type=1&id=659&version=2
func (m *AccountHealth) ShopPenalty() accounthealthentity.ShopPenaltyResult {
    method := "account_health/shop_penalty"
    result := accounthealthentity.ShopPenaltyResult{}
	params := lib.InRow{}
    err := m.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
