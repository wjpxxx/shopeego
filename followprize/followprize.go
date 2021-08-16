package followprize

import (
	"github.com/wjpxxx/letgo/lib"
	shopeeConfig "github.com/wjpxxx/shopeego/config"
	followprizeentity "github.com/wjpxxx/shopeego/followprize/entity"
)

//FollowPrize
type FollowPrize struct {
	Config *shopeeConfig.Config
}

//AddFollowPrize
//@Title OpenAPI add Follow Prize
//@Description https://open.shopee.com/documents?module=99&type=1&id=729&version=2
func (m *FollowPrize) AddFollowPrize(params followprizeentity.AddFollowPrizeRequest) followprizeentity.AddFollowPrizeResult {
    method := "follow_prize/add_follow_prize"
    result := followprizeentity.AddFollowPrizeResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//DeleteFollowPrize
//@Title delete_follow_prize
//@Description https://open.shopee.com/documents?module=99&type=1&id=730&version=2
func (m *FollowPrize) DeleteFollowPrize(params followprizeentity.DeleteFollowPrizeRequest) followprizeentity.DeleteFollowPrizeResult {
    method := "follow_prize/delete_follow_prize"
    result := followprizeentity.DeleteFollowPrizeResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//EndFollowPrize
//@Title end follow prize
//@Description https://open.shopee.com/documents?module=99&type=1&id=731&version=2
func (m *FollowPrize) EndFollowPrize(params followprizeentity.EndFollowPrizeRequest) followprizeentity.EndFollowPrizeResult {
    method := "follow_prize/end_follow_prize"
    result := followprizeentity.EndFollowPrizeResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//UpdateFollowPrize
//@Title update_follow_prize
//@Description https://open.shopee.com/documents?module=99&type=1&id=732&version=2
func (m *FollowPrize) UpdateFollowPrize(params followprizeentity.UpdateFollowPrizeRequest) followprizeentity.UpdateFollowPrizeResult {
    method := "follow_prize/update_follow_prize"
    result := followprizeentity.UpdateFollowPrizeResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//GetFollowPrizeDetail
//@Title get_follow_prize_detail
//@Description https://open.shopee.com/documents?module=99&type=1&id=733&version=2
func (m *FollowPrize) GetFollowPrizeDetail(params followprizeentity.GetFollowPrizeDetailRequest) followprizeentity.GetFollowPrizeDetailResult {
    method := "follow_prize/get_follow_prize_detail"
    result := followprizeentity.GetFollowPrizeDetailResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//GetFollowPrizeList
//@Title OpenAPI get_follow_prize_list
//@Description https://open.shopee.com/documents?module=99&type=1&id=734&version=2
func (m *FollowPrize) GetFollowPrizeList(params followprizeentity.GetFollowPrizeListRequest) followprizeentity.GetFollowPrizeListResult {
    method := "follow_prize/get_follow_prize_list"
    result := followprizeentity.GetFollowPrizeListResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
