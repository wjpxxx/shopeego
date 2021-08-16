package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetFollowPrizeListResult struct{
    Response	GetFollowPrizeListResponseResponseEntity	`json:"response"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
}
func (g GetFollowPrizeListResult) String() string {
    return lib.ObjectToString(g)
}

type GetFollowPrizeListResponseResponseEntity struct{
    More	bool	`json:"more"`
    FollowPrizeList	[]GetFollowPrizeListFollowPrizeListResponseEntity	`json:"follow_prize_list"`
}
func (g GetFollowPrizeListResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}