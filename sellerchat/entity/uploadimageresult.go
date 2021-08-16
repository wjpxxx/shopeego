package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UploadImageResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	UploadImageResponseResponseEntity	`json:"response"`
}
func (g UploadImageResult) String() string {
    return lib.ObjectToString(g)
}
