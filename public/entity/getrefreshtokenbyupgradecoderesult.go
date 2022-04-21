package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetRefreshTokenByUpgradeCodeResult struct{
    Response	GetRefreshTokenByUpgradeCodeResponseEntity	`json:"response"`
    RequestId	string	`json:"request_id"`
    More	bool	`json:"more"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
}
func (g GetRefreshTokenByUpgradeCodeResult) String() string {
    return lib.ObjectToString(g)
}
