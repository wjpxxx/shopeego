package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//AddGlobalModelResult
type AddGlobalModelResult struct {
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func (g AddGlobalModelResult) String() string {
	return lib.ObjectToString(g)
}
