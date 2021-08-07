package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//SearchItemResult
type SearchItemResult struct {
	commonentity.Result
	Response SearchItemResultResponse `json:"response"`
	Warning  []string                 `json:"warning"`
}

//String
func (r SearchItemResult) String() string {
	return lib.ObjectToString(r)
}

//SearchItemResultResponse
type SearchItemResultResponse struct {
	ItemIdList []int64 `json:"item_id_list"`
	TotalCount int     `json:"total_count"`
	NextOffset string  `json:"next_offset"`
}

//String
func (r SearchItemResultResponse) String() string {
	return lib.ObjectToString(r)
}
