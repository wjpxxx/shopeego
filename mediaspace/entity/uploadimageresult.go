package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//UploadImageResult
type UploadImageResult struct {
	commonentity.Result
	Error   string `json:"error"`
	Warning string `json:"warning"`
}

//String
func (g UploadImageResult) String() string {
	return lib.ObjectToString(g)
}
