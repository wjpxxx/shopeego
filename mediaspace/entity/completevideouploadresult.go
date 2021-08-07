package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//CompleteVideoUploadResult
type CompleteVideoUploadResult struct {
	commonentity.Result
	Error   string `json:"error"`
	Warning string `json:"warning"`
}

//String
func (g CompleteVideoUploadResult) String() string {
	return lib.ObjectToString(g)
}
