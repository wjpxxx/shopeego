package push

import (
	"github.com/wjpxxx/letgo/lib"
	shopeeConfig "github.com/wjpxxx/shopeego/config"
	pushentity "github.com/wjpxxx/shopeego/push/entity"
)

//Push
type Push struct {
	Config *shopeeConfig.Config
}

//GetPushConfig
//@Title get the configuration information of push service
//@Description https://open.shopee.com/documents?module=99&type=1&id=665&version=2
func (m *Push) GetPushConfig(params pushentity.GetPushConfigRequest) pushentity.GetPushConfigResult {
    method := "push/get_push_config"
    result := pushentity.GetPushConfigResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//SetPushConfig
//@Title Use this API to set the configuration information of push service.
//@Description https://open.shopee.com/documents?module=99&type=1&id=666&version=2
func (m *Push) SetPushConfig(params pushentity.SetPushConfigRequest) pushentity.SetPushConfigResult {
    method := "push/set_push_config"
    result := pushentity.SetPushConfigResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}

