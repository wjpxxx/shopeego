package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//CancelVideoUploadResult
type CancelVideoUploadResult struct {
	commonentity.Result
	Error   string `json:"error"`
	Warning string `json:"warning"`
}

//String
func (g CancelVideoUploadResult) String() string {
	return lib.ObjectToString(g)
}
