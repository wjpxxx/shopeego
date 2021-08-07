package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//UnSplitOrderResult
type AddInvoiceDataResult struct {
	commonentity.Result
}

//String
func (a AddInvoiceDataResult) String() string {
	return lib.ObjectToString(a)
}
