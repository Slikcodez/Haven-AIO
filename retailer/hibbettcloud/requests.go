package hibbettcloud

import (
	"main/client"
	"strings"

	http "github.com/bogdanfinn/fhttp"
)

func (self *HibbettBase) loginRequest(jsonData []byte) string {
	return client.TlsRequest(client.TLSParams{
		Client: self.client,
		Method: http.MethodPost,
		Url:    hibbettUrl + "login",
		Headers: http.Header{
			"Accept":             {"*/*"},
			"Accept-Encoding":    {"br;q=1.0, gzip;q=0.9, deflate;q=0.8"},
			"Accept-Language":    {"en-US;q=1.0"},
			"Connection":         {"keep-alive"},
			"Content-Type":       {"application/json; charset=utf-8"},
			"platform":           {"ios"},
			"version":            {"6.3.0"},
			"x-api-key":          {"0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0"},
			"X-PX-AUTHORIZATION": {"4"}, //1 also works
			"User-Agent":         {self.userAgent},
		},
		Body:             strings.NewReader(string(jsonData)),
		ExpectedResponse: 200,
	})
}

func (self *HibbettBase) paymentIDRequest() string {
	return client.TlsRequest(client.TLSParams{
		Client: self.client,
		Method: http.MethodGet,
		Url:    hibbettUrl + self.customerId + "/payment_methods",
		Headers: http.Header{
			"Accept":             {"*/*"},
			"Accept-Encoding":    {"br;q=1.0, gzip;q=0.9, deflate;q=0.8"},
			"Accept-Language":    {"en-US;q=1.0"},
			"Connection":         {"keep-alive"},
			"Content-Type":       {"application/json; charset=utf-8"},
			"platform":           {"ios"},
			"version":            {"6.3.0"},
			"Authorization":      {"Bearer " + self.sessionId},
			"x-api-key":          {"0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0"},
			"X-PX-AUTHORIZATION": {"4"}, //1 also works
			"User-Agent":         {self.userAgent},
		},
		Body:             nil,
		ExpectedResponse: 200,
	},
	)
}
