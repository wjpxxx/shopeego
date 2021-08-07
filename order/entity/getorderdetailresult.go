package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//GetOrderDetailResult
type GetOrderDetailResult struct {
	commonentity.Result
	Response GetOrderDetailResponse `json:"response"`
	Warning  []string               `json:"warning"`
}

//String
func (g GetOrderDetailResult) String() string {
	return lib.ObjectToString(g)
}

//GetOrderDetailResponse
type GetOrderDetailResponse struct {
	OrderList []OrderEntity `json:"order_list"`
}

//String
func (g GetOrderDetailResponse) String() string {
	return lib.ObjectToString(g)
}
