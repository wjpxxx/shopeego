package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//UpdateShippingOrderResult
type UpdateShippingOrderResult struct {
	commonentity.Result
}

//String
func (u UpdateShippingOrderResult) String() string {
	return lib.ObjectToString(u)
}
