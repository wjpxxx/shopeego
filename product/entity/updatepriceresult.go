package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//UpdatePriceResult
type UpdatePriceResult struct {
	commonentity.Result
	Warning  string                    `json:"warning"`
	Response UpdatePriceResultResponse `json:"response"`
}

//String
func (g UpdatePriceResult) String() string {
	return lib.ObjectToString(g)
}

//UpdatePriceResultResponse
type UpdatePriceResultResponse struct {
	FailureList []UpdatePriceFailureEntity `json:"failure_list"`
	SuccessList []UpdatePriceSuccessEntity `json:"success_list"`
}

//String
func (g UpdatePriceResultResponse) String() string {
	return lib.ObjectToString(g)
}
