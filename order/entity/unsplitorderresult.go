package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//UnSplitOrderResult
type UnSplitOrderResult struct {
	commonentity.Result
	Response OrderUpdateTimeResponse `json:"response"`
}

//String
func (u UnSplitOrderResult) String() string {
	return lib.ObjectToString(u)
}
