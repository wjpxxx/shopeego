package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//HandleBuyerCancellationResult
type HandleBuyerCancellationResult struct {
	commonentity.Result
}

//String
func (h HandleBuyerCancellationResult) String() string {
	return lib.ObjectToString(h)
}
