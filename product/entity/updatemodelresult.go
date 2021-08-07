package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//UpdateModelResult
type UpdateModelResult struct {
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func (g UpdateModelResult) String() string {
	return lib.ObjectToString(g)
}
