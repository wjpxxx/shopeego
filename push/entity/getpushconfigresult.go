package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetPushConfigResult struct{
    CallbackUrl	string	`json:"callback_url"`
    ShutTime	int	`json:"shut_time"`
    PushConfig	GetPushConfigPushConfigResponseEntity	`json:"push_config"`
    BlockedShopId	[]int64	`json:"blocked_shop_id"`
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
}
func (g GetPushConfigResult) String() string {
    return lib.ObjectToString(g)
}
