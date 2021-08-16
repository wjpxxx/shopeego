package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetMerchantsByPartnerResult struct{
    AuthedMerchantList	[]GetMerchantsByPartnerAuthedMerchantListResponseEntity	`json:"authed_merchant_list"`
    RequestId	string	`json:"request_id"`
    More	bool	`json:"more"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
}
func (g GetMerchantsByPartnerResult) String() string {
    return lib.ObjectToString(g)
}
