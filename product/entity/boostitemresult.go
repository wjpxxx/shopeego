package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//BoostItemResult
type BoostItemResult struct {
	commonentity.Result
	Warning  string                  `json:"warning"`
	Response BoostItemResultResponse `json:"response"`
}

//String
func (r BoostItemResult) String() string {
	return lib.ObjectToString(r)
}

//BoostItemResultResponse
type BoostItemResultResponse struct {
	FailureList []FailureEntity          `json:"failure_list"`
	SuccessList BoostItemSuccessEntity `json:"success_list"`
}

//String
func (r BoostItemResultResponse) String() string {
	return lib.ObjectToString(r)
}
