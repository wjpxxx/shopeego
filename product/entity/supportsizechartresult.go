package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//SupportSizeChartResult
type SupportSizeChartResult struct {
	commonentity.Result
	Warning  string                         `json:"warning"`
	Response SupportSizeChartResultResponse `json:"response"`
}

//String
func (r SupportSizeChartResult) String() string {
	return lib.ObjectToString(r)
}

//SupportSizeChartResultResponse
type SupportSizeChartResultResponse struct {
	SupportSizeChart bool `json:"support_size_chart"`
}

//String
func (r SupportSizeChartResultResponse) String() string {
	return lib.ObjectToString(r)
}
