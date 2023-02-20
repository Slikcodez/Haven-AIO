package hibbettcloud

import (
	"encoding/json"
	"fmt"
	"main/client"
	"main/constants"

	http "github.com/bogdanfinn/fhttp"
)

type Payment struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	PaymentObject struct {
		NameOnCard        string `json:"nameOnCard"`
		CardType          string `json:"cardType"`
		Number            string `json:"number"`
		ExpirationMonth   int    `json:"expirationMonth"`
		ExpirationYear    int    `json:"expirationYear"`
		CreditCardExpired bool   `json:"creditCardExpired"`
		CreditCardToken   string `json:"creditCardToken"`
		EncryptedCVNValue string `json:"encryptedCVNValue"`
	} `json:"paymentObject"`
}

func (user *HibbettBase) getPaymentId() (payments []Payment, err error) {
	res, err := client.TlsRequest(client.TLSParams{
		Client: user.client,
		Method: http.MethodGet,
		Url:    constants.GetPaymentIdUrlString(user.customerId),
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
			"X-PX-AUTHORIZATION": {"4"}, //1 also works
			"User-Agent":         {user.userAgent},
		},
		Body:             nil,
		ExpectedResponse: 200,
	},
	)
	if err != nil {
		fmt.Println("Error getting payment id")
		Init(user.thread, user.account)
		return
	}
	payments = user.unmarshalPaymentIDs(res)
	return
}

func (user *HibbettBase) unmarshalPaymentIDs(payload []byte) (payments []Payment) {

	err := json.Unmarshal([]byte(payload), &payments)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
		// Init(self.thread, self.account)
	} else {

		for _, payment := range payments {

			if payment.PaymentObject.Number == user.four {
				user.paymentId = payment.ID
				user.paymentType = payment.Type
				fmt.Println("Thread " + user.thread + ": Got PaymentID")
			} else {
				fmt.Println("Error, no valid paymentID found")
			}
		}

	}
	return
}
