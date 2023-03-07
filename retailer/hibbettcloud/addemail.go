package hibbettcloud

import (
	"encoding/json"
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
		user.loginAccount()

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
			"Accept-Language":     {"es-US;q=0.9"},
			"Connection":          {"keep-alive"},
			"Content-Type":        {"application/json; charset=utf-8"},
			"platform":            {"ios"},
			"version":             {"6.3.0"},
			"Authorization":       {"Bearer " + user.sessionId},
			"x-api-key":           {"0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0"},
			"X-PX-AUTHORIZATION":  {"2"}, //1 also works
			"X-PX-ORIGINAL-TOKEN": {"2:" + constants.RandString()},
			"Cache-Control":       {"max-age=0"},
			"User-Agent":          {user.userAgent},
		},
		Body:             strings.NewReader(string(jsonData)),
		ExpectedResponse: 200,
	},
	)

	return
}
