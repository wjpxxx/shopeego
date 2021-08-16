package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DisputeRequest struct{
    ReturnSn	string	`json:"return_sn"`
    Email	string	`json:"email"`
    DisputeReason	string	`json:"dispute_reason"`
    DisputeTextReason	string	`json:"dispute_text_reason"`
    Image	[]string	`json:"image"`
}
func (g DisputeRequest) String() string {
    return lib.ObjectToString(g)
}
