package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//GetOrderListResult
type GetOrderListResult struct {
	commonentity.Result
	Response GetOrderListResponse `json:"response"`
}

//String
func (g GetOrderListResult) String() string {
	return lib.ObjectToString(g)
}

//GetOrderListResponse
type GetOrderListResponse struct {
	More       bool          `json:"more"`
	NextCursor string        `json:"next_cursor"`
	OrderList  []OrderEntity `json:"order_list"`
}

//String
func (g GetOrderListResponse) String() string {
	return lib.ObjectToString(g)
}
