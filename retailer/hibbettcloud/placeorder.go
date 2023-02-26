package hibbettcloud

import (
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	"main/client"
	"main/constants"
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
			} else {
				constants.LogStatus(user.thread, "Declined")
			}
			constants.LogStatus(user.thread, "Invalid CVV")

			fmt.Println(constants.UnmarshalRequestError(err.Error(), "body"))
		}
	} else {
		fmt.Println(res)
	}

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
