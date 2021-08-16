package sellerchat

import (
	"github.com/wjpxxx/letgo/lib"
	shopeeConfig "github.com/wjpxxx/shopeego/config"
	sellerchatentity "github.com/wjpxxx/shopeego/sellerchat/entity"
)

//Sellerchat
type Sellerchat struct {
	Config *shopeeConfig.Config
}

//GetMessage
//@Title To get messages history for a specific conversation, which can display the messages detail from sender and receiver.
//@Description https://open.shopee.com/documents?module=99&type=1&id=671&version=2
func (m *Sellerchat) GetMessage(params sellerchatentity.GetMessageRequest) sellerchatentity.GetMessageResult {
    method := "sellerchat/get_message"
    result := sellerchatentity.GetMessageResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//SendMessage
//@Title 1.To send a message and select the correct message type (Do not use this API to send batch messages) 
//2.Currently TW region is not supported to send messages.
//@Description https://open.shopee.com/documents?module=99&type=1&id=672&version=2
func (m *Sellerchat) SendMessage(params sellerchatentity.SendMessageRequest) sellerchatentity.SendMessageResult {
    method := "sellerchat/send_message"
    result := sellerchatentity.SendMessageResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//GetConversationList
//@Title To get conversation list and its params data
//@Description https://open.shopee.com/documents?module=99&type=1&id=673&version=2
func (m *Sellerchat) GetConversationList(params sellerchatentity.GetConversationListRequest) sellerchatentity.GetConversationListResult {
    method := "sellerchat/get_conversation_list"
    result := sellerchatentity.GetConversationListResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//GetOneConversation
//@Title To get a specific conversation's basic information.
//@Description https://open.shopee.com/documents?module=99&type=1&id=674&version=2
func (m *Sellerchat) GetOneConversation(params sellerchatentity.GetOneConversationRequest) sellerchatentity.GetOneConversationResult {
    method := "sellerchat/get_one_conversation"
    result := sellerchatentity.GetOneConversationResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//DeleteConversation
//@Title To delete a specific conversation
//@Description https://open.shopee.com/documents?module=99&type=1&id=675&version=2
func (m *Sellerchat) DeleteConversation(params sellerchatentity.DeleteConversationRequest) sellerchatentity.DeleteConversationResult {
    method := "sellerchat/delete_conversation"
    result := sellerchatentity.DeleteConversationResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//PinConversation
//@Title To pin a specific conversation
//@Description https://open.shopee.com/documents?module=99&type=1&id=677&version=2
func (m *Sellerchat) PinConversation(params sellerchatentity.PinConversationRequest) sellerchatentity.PinConversationResult {
    method := "sellerchat/pin_conversation"
    result := sellerchatentity.PinConversationResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//UnpinConversation
//@Title To unpin a specific conversation
//@Description https://open.shopee.com/documents?module=99&type=1&id=678&version=2
func (m *Sellerchat) UnpinConversation(params sellerchatentity.UnpinConversationRequest) sellerchatentity.UnpinConversationResult {
    method := "sellerchat/unpin_conversation"
    result := sellerchatentity.UnpinConversationResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//ReadConversation
//@Title To send read request for a specific conversation
//@Description https://open.shopee.com/documents?module=99&type=1&id=679&version=2
func (m *Sellerchat) ReadConversation(params sellerchatentity.ReadConversationRequest) sellerchatentity.ReadConversationResult {
    method := "sellerchat/read_conversation"
    result := sellerchatentity.ReadConversationResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//UnreadConversation
//@Title To mark a conversation as unread
//@Description https://open.shopee.com/documents?module=99&type=1&id=680&version=2
func (m *Sellerchat) UnreadConversation(params sellerchatentity.UnreadConversationRequest) sellerchatentity.UnreadConversationResult {
    method := "sellerchat/unread_conversation"
    result := sellerchatentity.UnreadConversationResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//GetOfferToggleStatus
//@Title To get the toggle status to check if the shop has allowed buyer to negotiate price with seller.
//@Description https://open.shopee.com/documents?module=99&type=1&id=681&version=2
func (m *Sellerchat) GetOfferToggleStatus(params sellerchatentity.GetOfferToggleStatusRequest) sellerchatentity.GetOfferToggleStatusResult {
    method := "sellerchat/get_offer_toggle_status"
    result := sellerchatentity.GetOfferToggleStatusResult{}
    err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//SetOfferToggleStatus
//@Title To set the toggle status.If set as "enabled", then seller doesn't allow buyer negotiate the price.
//@Description https://open.shopee.com/documents?module=99&type=1&id=682&version=2
func (m *Sellerchat) SetOfferToggleStatus(params sellerchatentity.SetOfferToggleStatusRequest) sellerchatentity.SetOfferToggleStatusResult {
    method := "sellerchat/set_offer_toggle_status"
    result := sellerchatentity.SetOfferToggleStatusResult{}
    err := m.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
//SUploadImage
//@Title When you need to send an image type message, please request this API first to upload the image file to get image url. Then proceed to request the send message API with the image url.
//@Description https://open.shopee.com/documents?module=99&type=1&id=683&version=2
func (m *Sellerchat) SUploadImage(file string) sellerchatentity.UploadImageResult {
    method := "sellerchat/upload_image"
    result := sellerchatentity.UploadImageResult{}
    params := lib.InRow{
		"@file": file,
	}
    err := m.Config.HttpPostFile(method, params, &result)
    if err != nil {
        result.Error = err.Error()
    }
    return result
}
