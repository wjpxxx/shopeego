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
	Response Response `json:"response"`
}

//String
func (g UploadImageResult) String() string {
	return lib.ObjectToString(g)
}
//Response
type Response struct{
	ImageInfo ImageInfo `json:"image_info"`
}

//String
func (g Response) String() string {
	return lib.ObjectToString(g)
}

//ImageInfo
type ImageInfo struct{
	ImageId string `json:"image_id"`
	ImageUrlList []ImageUrlList `json:"image_url_list"`
}

//String
func (g ImageInfo) String() string {
	return lib.ObjectToString(g)
}

type ImageUrlList struct{
	ImageUrlRegion string `json:"image_url_region"`
	ImageUrl string `json:"image_url"`
}

//String
func (g ImageUrlList) String() string {
	return lib.ObjectToString(g)
}