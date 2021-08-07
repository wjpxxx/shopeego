package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//GetItemListResult
type GetGlobalItemListResult struct {
	commonentity.Result
	Response GetGlobalItemListResultResponse `json:"response"`
	Warning  string                          `json:"warning"`
}

//String
func (g GetGlobalItemListResult) String() string {
	return lib.ObjectToString(g)
}

//GetItemListResultResponse
type GetGlobalItemListResultResponse struct {
	GlobalItemList []GlobalItemListEntity `json:"global_item_list"`
	TotalCount     int                    `json:"total_count"`
	HasNextPage    bool                   `json:"has_next_page"`
	Offset         string                 `json:"offset"`
}

//String
func (g GetGlobalItemListResultResponse) String() string {
	return lib.ObjectToString(g)
}
