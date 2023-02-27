package hibbettcloud

import (
	"encoding/json"
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	"main/client"
	"main/constants"
	webhook "main/webhooks"
	"strings"
)

func (user *HibbettBase) placeOrder() {

	constants.LogStatus(user.thread, "Placing Order")

	res, err := user.placeOrderRequest()

	if err != nil {
		StatusCode := constants.UnmarshalRequestError(err.Error(), "status")

		if StatusCode == "403" {
			constants.LogStatus(user.thread, "PX Blocked While Placing Order")
		}
		if StatusCode == "400" {
			if strings.Contains("Invalid", constants.UnmarshalRequestError(err.Error(), "body")) {
				constants.LogStatus(user.thread, "Invalid Card Cvv")
				constants.Declines++
			} else {
				constants.LogStatus(user.thread, "Declined")
				constants.Declines++
			}

		}
	} else {
		var Order Order
		if err := json.Unmarshal(res, &Order); err != nil {
			panic(err)
		}
		
		err := webhook.SendWebhook(Order.OrderItems[0].Sku.Size, Order.OrderItems[0].MasterID, Order.Total, Order.ID, Order.OrderItems[0].Product.ImageResources["0001-0"][0].URL, user.email)
		if err != nil {
			return
		}
	}

}

type Order struct {
	Adjustments []struct{} `json:"adjustments"`
	ID          string     `json:"id"`
	OrderItems  []struct {
		MasterID string `json:"masterId"`
		Product  struct {
			ImageResources map[string][]struct {
				URL   string `json:"url"`
				Usage string `json:"usage"`
			} `json:"imageResources"`
			Name string `json:"name"`
		} `json:"product"`
		Sku struct {
			AvailableQuantity interface{} `json:"availableQuantity"`
			Color             struct {
				ID           string `json:"id"`
				ImagePattern struct {
					URL   string `json:"url"`
					Usage string `json:"usage"`
				} `json:"imagePattern"`
				Label string `json:"label"`
			} `json:"color"`
			DiscountedPrice interface{} `json:"discountedPrice"`
			FinalPrice      string      `json:"finalPrice"`
			ID              string      `json:"id"`
			ListPrice       string      `json:"listPrice"`
			ProductID       string      `json:"productId"`
			SelectedOptions []struct{}  `json:"selectedOptions"`
			Size            string      `json:"size"`
			SkuNumber       string      `json:"skuNumber"`
		} `json:"sku"`
	} `json:"orderItems"`
	Total float64 `json:"total"`
}

func (user *HibbettBase) placeOrderRequest() (res []byte, err error) {
	res, err = client.TlsRequest(client.TLSParams{
		Client: user.client,
		Method: http.MethodPost,
		Url:    fmt.Sprintf(`https://hibbett-mobileapi.prolific.io/ecommerce/cart/%s/place_order?cardSecurityCode=%s&customer=%s&phone=&oneTapCheckout=true&firstName=&optIn=false`, user.cartId, user.cvv, user.customerId),
		Headers: http.Header{
			"Accept":             {"*/*"},
			"Accept-Encoding":    {"br;q=1.0, gzip;q=0.9, deflate;q=0.8"},
			"Accept-Language":    {"en-US;q=1.0"},
			"Connection":         {"keep-alive"},
			"Content-Type":       {"application/json; charset=utf-8"},
			"platform":           {"ios"},
			"version":            {"6.3.0"},
			"Authorization":      {"Bearer " + user.sessionId},
			"x-api-key":          {"0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0"},
			"X-PX-AUTHORIZATION": {"2"}, //1 also works
			"User-Agent":         {user.userAgent},
		},
		Body:             nil,
		ExpectedResponse: 200,
	},
	)

	return
}
