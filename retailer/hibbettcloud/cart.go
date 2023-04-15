package hibbettcloud

import (
	"encoding/json"
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	"main/client"
	"main/constants"
	"strings"
	"time"
)

type PreCart struct {
	PreferredBillingAddressId  string `json:"preferredBillingAddressId"`
	PreferredShippingAddressId string `json:"preferredShippingAddressId"`
	CartItems                  []struct {
		Quantity         int           `json:"quantity"`
		Personalizations []interface{} `json:"personalizations"`
		Product          struct {
			Id string `json:"id"`
		} `json:"product"`
		CustomerId string `json:"customerId"`
		Sku        struct {
			Id string `json:"id"`
		} `json:"sku"`
	} `json:"cartItems"`
	CustomerId               string `json:"customerId"`
	PreferredPaymentMethodId string `json:"preferredPaymentMethodId"`
}

type PreCartRes struct {
	CartID    string `json:"id"`
	SessionID string `json:"bmSessionToken"`
}

func (user *HibbettBase) unmarshalPreCart(res []byte, sku string) (err error, precart PreCartRes) {

	err = json.Unmarshal([]byte(res), &precart)
	if err != nil {
		fmt.Println(err)
		constants.LogStatus(user.thread, "Error Parsing JSON")
		fmt.Println(err)
		Init(user.thread, user.account, user.mode)
	} else {
		user.cartId = precart.CartID
		user.sessionId = precart.SessionID
		constants.LogStatus(user.thread, "Carted "+sku)
		constants.Carts++
		user.addEmail()

	}

	return
}

func (user *HibbettBase) preCart(productInfo string) {

	sku := strings.Split(productInfo, ":")[0]

	precart := &PreCart{
		PreferredBillingAddressId:  "main",
		PreferredShippingAddressId: "main",
		CartItems: []struct {
			Quantity         int           `json:"quantity"`
			Personalizations []interface{} `json:"personalizations"`
			Product          struct {
				Id string `json:"id"`
			} `json:"product"`
			CustomerId string `json:"customerId"`
			Sku        struct {
				Id string `json:"id"`
			} `json:"sku"`
		}{
			{
				Quantity:         1,
				Personalizations: []interface{}{},
				Product: struct {
					Id string `json:"id"`
				}{
					Id: sku,
				},
				CustomerId: user.customerId,
				Sku: struct {
					Id string `json:"id"`
				}{
					Id: sku,
				},
			},
		},
		CustomerId:               user.customerId,
		PreferredPaymentMethodId: user.paymentId,
	}
	jsonData, _ := json.Marshal(precart)

	constants.LogStatus(user.thread, "Initializing Cart")
	res, err123 := user.preCartRequest(jsonData)
	if err123 != nil {

		status_code, err11 := constants.UnmarshalRequestError(err123.Error(), "status")
		body, _ := constants.UnmarshalRequestError(err123.Error(), "body")
		if err11 != nil {
			constants.LogStatus(user.thread, "ERROR AT CART")
			fmt.Println(err11)
			user.loginAccount()
		} else {
			if status_code == "403" {
				constants.LogStatus(user.thread, "PX Blocked While Carting")

				if time.Now().Second() < 15 {
					user.getProxy()
					user.preCart(productInfo)
				} else {
					user.loginAccount()
				}
			}
			if status_code == "400" {
				if strings.Contains(body, "One Tap") {
					constants.LogStatus(user.thread, "Unable to create cart")
				} else {
					constants.LogStatus(user.thread, body)
				}
			}
			Init(user.thread, user.account, user.mode)
		}
	} else {
		user.unmarshalPreCart(res, sku)
	}

}

func (user *HibbettBase) preCartRequest(jsonData []byte) (res []byte, err error) {

	pxvals := make([]string, 0)
	pxvals = append(pxvals, "2")
	//pxvals = append(pxvals, "4")

	res, err = client.TlsRequest(client.TLSParams{
		Client: user.client,
		Method: http.MethodPost,
		Url:    `https://hibbett-mobileapi.prolific.io/ecommerce/cart/one_tap?cardSecurityCode=` + user.cvv,
		Headers: http.Header{
			"Accept":             {"*/*"},
			"Accept-Encoding":    {"br;q=1.0, gzip;q=0.9, deflate;q=0.8"},
			"Accept-Language":    {"es-US;q=0.9"},
			"Connection":         {"keep-alive"},
			"Content-Type":       {"application/json; charset=utf-8"},
			"platform":           {"ios"},
			"version":            {"6.3.2"},
			"Authorization":      {"Bearer " + user.sessionId},
			"x-api-key":          {"0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0"},
			"X-PX-AUTHORIZATION": {"3:" + user.pxToken}, //1 also works
			"User-Agent":         {user.userAgent},
		},
		Body:             strings.NewReader(string(jsonData)),
		ExpectedResponse: 200,
	},
	)

	return
}
