package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//InitTierVariationResult
type InitTierVariationResult struct {
	commonentity.Result
	Response InitTierVariationResultResponse `json:"response"`
	Warning  string                          `json:"warning"`
}

//String
func (g InitTierVariationResult) String() string {
	return lib.ObjectToString(g)
}

//GetModelListResultResponse
type InitTierVariationResultResponse struct {
	ItemID        int64                 `json:"item_id"`
	TierVariation []TierVariationEntity `json:"tier_variation"`
	Model         []ModelEntity         `json:"model"`
}

//String
func (g InitTierVariationResultResponse) String() string {
	return lib.ObjectToString(g)
}
