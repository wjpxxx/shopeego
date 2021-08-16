package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetOfferToggleStatusResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	SetOfferToggleStatusResponseResponseEntity	`json:"response"`
}
func (g SetOfferToggleStatusResult) String() string {
    return lib.ObjectToString(g)
}


type SetOfferToggleStatusResponseResponseEntity struct{

}

func (g SetOfferToggleStatusResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}