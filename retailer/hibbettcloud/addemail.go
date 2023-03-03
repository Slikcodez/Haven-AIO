package hibbettcloud

import (
	"encoding/json"
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	"main/client"
	"main/constants"
	"strings"
)

type email struct {
	Email string `json:"email"`
}

func (user *HibbettBase) addEmail() {
	addemail := &email{
		Email: user.email,
	}
	jsondata, _ := json.Marshal(addemail)
	_, err := user.addEmailRequest(jsondata)
	if err != nil {
		joe, _ := constants.UnmarshalRequestError(err.Error(), "body")
		fmt.Println(joe)
	} else {
		constants.LogStatus(user.thread, "Added Order Information")
		user.placeOrder()
	}

}

func (user *HibbettBase) addEmailRequest(jsonData []byte) (res []byte, err error) {

	res, err = client.TlsRequest(client.TLSParams{
		Client: user.client,
		Method: http.MethodPut,
		Url:    `https://hibbett-mobileapi.prolific.io/ecommerce/cart/` + user.cartId + `/customer`,
		Headers: http.Header{
			"Accept":              {"*/*"},
			"Accept-Encoding":     {"br;q=1.0, gzip;q=0.9, deflate;q=0.8"},
			"Accept-Language":     {"en-US;q=1.0"},
			"Connection":          {"keep-alive"},
			"Content-Type":        {"application/json; charset=utf-8"},
			"platform":            {"ios"},
			"version":             {"6.3.0"},
			"x-api-key":           {"0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0"},
			"X-PX-AUTHORIZATION":  {"2"},
			"X-PX-ORIGINAL-TOKEN": {"2:" + constants.RandString()}, //1 also works
			"User-Agent":          {user.userAgent},
		},
		Body:             strings.NewReader(string(jsonData)),
		ExpectedResponse: 200,
	},
	)

	return
}
