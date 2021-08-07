package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//SetNoteResult
type SetNoteResult struct {
	commonentity.Result
}

//String
func (s SetNoteResult) String() string {
	return lib.ObjectToString(s)
}
