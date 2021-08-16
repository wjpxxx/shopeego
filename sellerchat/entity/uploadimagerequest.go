package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UploadImageRequest struct{
    File	string	`json:"file"`  //文件路径
}
func (g UploadImageRequest) String() string {
    return lib.ObjectToString(g)
}

