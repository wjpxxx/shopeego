package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetEscrowListResponseResponseEntity struct{
    EscrowList	[]GetEscrowListEscrowListResponseEntity	`json:"escrow_list"`
    More	bool	`json:"more"`
}
func (g GetEscrowListResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
