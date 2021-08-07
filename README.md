# shopee go language API

API interface for shopee go language version

my email :474790700@qq.com

## Contents

- [shopee go language API](#shopee-go)
  - [Installation](#installation)
  - [Quick start](#quick-start)

## Installation

To install shopeego package, you need to install Go and set your Go workspace first.

1. The first need [Go](https://golang.org/) installed (**version 1.12+ is required**), then you can use the below Go command to install shopeego.

```sh
$ go get -u github.com/wjpxxx/shopeego
```

2. Import it in your code:

```go
import (
	"github.com/wjpxxx/shopeego"
)
```
## Quick start

## API call

```go
package main

import (
	"github.com/wjpxxx/shopeego"
	"github.com/wjpxxx/shopeego/commonentity"
	shopeeConfig "github.com/wjpxxx/shopeego/config"

	//"github.com/wjpxxx/shopeego/order"
	//orderEntity "github.com/wjpxxx/shopeego/order/entity"
	//"github.com/wjpxxx/letgo/file"
	//"github.com/wjpxxx/letgo/lib"
	"fmt"
)

func main() {
	shopeego.Register("shopee-api", shopeeConfig.New("https://partner.test-stable.shopeemobile.com", "/api/v2/", 1001219, "cea778f3b36d99bda5d16a4e511fa55f9032464940163fe4acfee13c48658f42", "/shopee_callback"))
	//fmt.Println(shopeego.GetApi("shopee-api").AuthorizationURL())
	//file.PutContent("json",fmt.Sprintf("%v",shopeego.GetApi("shopee-api").GetAccesstoken("4f9a5a2b011ce202f7f5e51db98d5024",9714)))
	//fmt.Println(shopeego.GetApi("shopee-api").RefreshAccessToken(*commonentity.NewShop(9714,14377,"9938a79551c3e6463b2c5b92997a2298","5e02b201225286bb960fc47e4e3f7554")))
	/*
		shopeego.Register("shopee-api-v2",shopeeConfig.New("https://partner.test-stable.shopeemobile.com","/api/v2/",1001219,"cea778f3b36d99bda5d16a4e511fa55f9032464940163fe4acfee13c48658f42","/shopee_callback").SetShopInfo(&commonentity.ShopInfo{
			RefreshToken:"9a9468cf1d0f8e972765cbe779fe6b45",
			AccessToken:"e69f7c8a64622c5b02e6166d1ee19877",
			ExpireIn:14376,
			ShopID:9714,
		}))
	*/
	//shopeego.GetApi("shopee-api-v2").GetOrderList("create_time",lib.Time()-3600*24*10,lib.Time(),20,"",order.UNPAID,"order_status")
	//shopeego.GetApi("shopee-api-v2").GetOrderDetail([]string{"210606JQ3AFK4A"})
	/*
		shopeego.GetApi("shopee-api-v2").SplitOrder("210606JQ3AFK4A",[]orderEntity.PackageListRequestEntity{
			orderEntity.PackageListRequestEntity{
				ItemList:[]orderEntity.PackageListRequestItemListEntity{
					orderEntity.PackageListRequestItemListEntity{
						ItemID:100015844,
					},
				},
			},
		})
		shopeego.GetApi("shopee-api-v2").CancelOrder("210606JQ3AFK4A",order.OUT_OF_STOCK,[]orderEntity.CancelOrderRequestEntity{
			orderEntity.CancelOrderRequestEntity{
				ItemID:100015844,
				ModelID: 10000083295,
			},
		})
		shopeego.GetApi("shopee-api-v2").GetShippingParameter("210606JQ3AFK4A")
		//fmt.Println(shopeego.GetApi("shopee-api-v2").GetItemBaseInfo([]int64{100015844}))
		//fmt.Println(shopeego.GetApi("shopee-api-v2").GetItemExtraInfo([]int64{100015844}))
		shopeego.GetApi("shopee-api-v2").GetItemList(0,10,0,0,"NORMAL")
	*/
	shopeego.Register("shopee-api-v2", shopeeConfig.New("https://partner.test-stable.shopeemobile.com", "/api/v2/", 1001219, "cea778f3b36d99bda5d16a4e511fa55f9032464940163fe4acfee13c48658f42", "/shopee_callback").SetShopInfo(&commonentity.ShopInfo{
		RefreshToken: "2a72d96edeab0388fba787ce71031b5b",
		AccessToken:  "623f3183d63c5cd6f8c5f27e060dc2ec",
		ExpireIn:     14376,
		ShopID:       9714,
	}))
	fmt.Println(shopeego.GetApi("shopee-api-v2").GetMerchantInfo())
}

```
