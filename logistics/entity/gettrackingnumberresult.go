package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//GetTrackingNumberResult
type GetTrackingNumberResult struct {
	commonentity.Result
	Response TrackingNumberEntity `json:"response"`
}

//String
func (g GetTrackingNumberResult) String() string {
	return lib.ObjectToString(g)
}
