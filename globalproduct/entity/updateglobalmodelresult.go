package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//UpdateGlobalModelResult
type UpdateGlobalModelResult struct {
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func (g UpdateGlobalModelResult) String() string {
	return lib.ObjectToString(g)
}
