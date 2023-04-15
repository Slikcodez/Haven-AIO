package nordstrom

import (
	"encoding/json"
	"github.com/google/uuid"
	"main/client"
	"main/constants"
	"strings"

	htttp "github.com/bogdanfinn/fhttp"
)

type Item struct {
	RmsSku   string `json:"rmsSku"`
	Quantity int    `json:"quantity"`
	StyleID  string `json:"styleId"`
}

func (user *NordstromBase) Cart() error {
	item := Item{
		RmsSku:   strings.Split(user.sku, ":")[0],
		Quantity: 1,
		StyleID:  strings.Split(user.sku, ":")[1],
	}

	jsonData, err := json.Marshal(item)
	if err != nil {
		constants.LogNordstromStatus(user.thread, "Error While Carting")
	}

	_, err69 := user.CartRequest(jsonData)
	if err69 != nil {
		constants.LogNordstromStatus(user.thread, "Blocked While Carting")
	} else {
		user.DoPayemnt()
	}

}

func (user *NordstromBase) CartRequest(jsonData []byte) (res []byte, err error) {

	header, err12 := getHeadersReq()

	if err12 != nil {
		constants.LogNordstromStatus(user.thread, "Error Getting Shape While Carting")
		return nil, err12
	}
	tracecontext := uuid.New()
	res, err = client.TlsRequest(client.TLSParams{
		Client: user.client,
		Method: htttp.MethodPost,
		Url:    "constants.HibbettURL" + user.profile.BillingAddress.PostCode,
		Headers: htttp.Header{
			"accept":            {"application/vnd.nord.checkout.v2+json"},
			"accept-language":   {"en-US,en;q=0.9"},
			"content-type":      {"application/json"},
			"identified-bot":    {"false"},
			"nord-country-code": {"US"},
			"sec-fetch-dest":    {"empty"},
			"sec-fetch-mode":    {"cors"},
			"sec-fetch-site":    {"same-origin"},
			"tracecontext":      {tracecontext.String()},
			"x-forter-token":    {strings.Split(header, ":::")[0]},
			"x-nor-appiid":      {strings.Split(header, ":::")[1]},
			"x-shopper-id":      {user.shopperID},
			"x-shopper-token":   {user.shopperToken},
			"X-y8S6k3DB-f":      {strings.Split(header, ":::")[2]},
			"X-y8S6k3DB-d":      {strings.Split(header, ":::")[3]},
			"X-y8S6k3DB-b":      {strings.Split(header, ":::")[4]},
			"X-y8S6k3DB-c":      {strings.Split(header, ":::")[5]},
			"X-y8S6k3DB-a":      {strings.Split(header, ":::")[6]},
			"X-y8S6k3DB-z":      {strings.Split(header, ":::")[7]},
		},
		Body:             strings.NewReader(string(jsonData)),
		ExpectedResponse: 200,
	})

	return
}
