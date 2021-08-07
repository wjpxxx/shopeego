package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//ShipOrderResult
type ShipOrderResult struct {
	commonentity.Result
}

//String
func (s ShipOrderResult) String() string {
	return lib.ObjectToString(s)
}
