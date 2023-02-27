package hibbettcloud

import (
	"encoding/json"
	"main/client"
	"main/constants"
	"strings"
	"time"

	http "github.com/bogdanfinn/fhttp"
)

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Session struct {
	SessionID  string `json:"sessionID"`
	CustomerID string `json:"customerID"`
}

/*
Function to login to account
*/
func (user *HibbettBase) loginAccount() {

	creds := &Credentials{
		Login:    user.email,
		Password: user.password,
	}

	jsonData, err := json.Marshal(creds)
	if err != nil {
		constants.LogStatus(user.thread, "Error Parsing JSON")
		return
	}

	res, err := user.loginRequest(jsonData)
	if err != nil {
		constants.LogStatus(user.thread, "Error Logging In")
		time.Sleep(10 * time.Second)
		Init(user.thread, user.account)
	}
	var responseData Session
	err = json.Unmarshal([]byte(res), &responseData)
	if err != nil {
		constants.LogStatus(user.thread, "Error Parsing JSON")
		Init(user.thread, user.account)
	} else {
		user.sessionId = responseData.SessionID
		user.customerId = responseData.CustomerID

		constants.LogStatus(user.thread, "Logged In Successfully")
		user.getPaymentId()
	}

}

func (user *HibbettBase) loginRequest(jsonData []byte) (res []byte, err error) {

	constants.LogStatus(user.thread, "Logging In")

	res, err = client.TlsRequest(client.TLSParams{
		Client: user.client,
		Method: http.MethodPost,
		Url:    constants.HibbettURL + "login",
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
			"X-PX-ORIGINAL-TOKEN": {"2:14265r7t638yiuhojnlkm;afs"}, //1 also works
			"User-Agent":          {user.userAgent},
		},
		Body:             strings.NewReader(string(jsonData)),
		ExpectedResponse: 200,
	},
	)
	return
}
