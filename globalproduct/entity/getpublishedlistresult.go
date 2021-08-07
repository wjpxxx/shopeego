package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//GetPublishedListResult
type GetPublishedListResult struct {
	commonentity.Result
	Response GetPublishedListResultResponse `json:"response"`
	Warning  string                         `json:"warning"`
}

//String
func (g GetPublishedListResult) String() string {
	return lib.ObjectToString(g)
}

//GetPublishedListResultResponse
type GetPublishedListResultResponse struct {
	PublishedItem []PublishedItemEntity `json:"published_item"`
}

//String
func (g GetPublishedListResultResponse) String() string {
	return lib.ObjectToString(g)
}
