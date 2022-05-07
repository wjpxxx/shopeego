package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//AddItemResult
type AddItemResult struct {
	commonentity.Result
	Warning       string     `json:"warning"`
	ItemDangerous int        `json:"item_dangerous"`
	Response      AddItemItemEntity `json:"response"`
}

//String
func (g AddItemResult) String() string {
	return lib.ObjectToString(g)
}
