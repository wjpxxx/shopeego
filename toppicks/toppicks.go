package toppicks

import (
	"github.com/wjpxxx/letgo/lib"
	shopeeConfig "github.com/wjpxxx/shopeego/config"
	toppicksentity "github.com/wjpxxx/shopeego/toppicks/entity"
)

//TopPicks
type TopPicks struct {
	Config *shopeeConfig.Config
}

//GetTopPicksList
//@Title get one TopPicks
//@Description https://open.shopee.com/documents?module=99&type=1&id=578&version=2
func (m *TopPicks) GetTopPicksList(params toppicksentity.GetTopPicksListRequest) toppicksentity.GetTopPicksListResult {
    method := "top_picks/get_top_picks_list"
    result := toppicksentity.GetTopPicksListResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//AddTopPicks
//@Title add one collection 
//@Description https://open.shopee.com/documents?module=99&type=1&id=579&version=2
func (m *TopPicks) AddTopPicks(params toppicksentity.AddTopPicksRequest) toppicksentity.AddTopPicksResult {
    method := "top_picks/add_top_picks"
    result := toppicksentity.AddTopPicksResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//UpdateTopPicks
//@Title update a collection info
//@Description https://open.shopee.com/documents?module=99&type=1&id=580&version=2
func (m *TopPicks) UpdateTopPicks(params toppicksentity.UpdateTopPicksRequest) toppicksentity.UpdateTopPicksResult {
    method := "top_picks/update_top_picks"
    result := toppicksentity.UpdateTopPicksResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//DeleteTopPicks
//@Title delete a collection
//@Description https://open.shopee.com/documents?module=99&type=1&id=581&version=2
func (m *TopPicks) DeleteTopPicks(params toppicksentity.DeleteTopPicksRequest) toppicksentity.DeleteTopPicksResult {
    method := "top_picks/delete_top_picks"
    result := toppicksentity.DeleteTopPicksResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
