package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//UpdateGlobalPriceResult
type UpdateGlobalPriceResult struct {
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func (g UpdateGlobalPriceResult) String() string {
	return lib.ObjectToString(g)
}
