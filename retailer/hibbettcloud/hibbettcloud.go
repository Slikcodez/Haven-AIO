package hibbettcloud

import (
	"bufio"
	"fmt"
	"main/client"
	"math/rand"
	"os"
	"strings"
	"time"
)

type HibbettBase struct {
	client      client.HttpClient
	thread      string
	account     string
	paymentId   string
	token       string
	userAgent   string
	paymentType string
	sessionId   string
	sessionEX   string
	customerId  string
	email       string
	password    string
	four        string
	checkedOut  bool
	genningPx   bool
	proxy       string
	initialSku  string
	detected    bool
	sku         string
	cartId      string
	cvv         string
}

func Init(thread string, account string) {

	user := HibbettBase{
		thread:      thread,
		account:     account,
		paymentId:   "",
		token:       "",
		userAgent:   fmt.Sprintf("hibbett | CG/6.3.0 (com.hibbett.hibbett-sports; build:%v; iOS 16.0.0)", rand.Intn(15000)+1),
		paymentType: "",
		sessionId:   "",
		sessionEX:   "",
		customerId:  "",
		email:       strings.Split(account, ":")[0],
		password:    strings.Split(account, ":")[1],
		four:        strings.Split(account, ":")[2],
		cvv:         strings.Split(account, ":")[3],
		checkedOut:  false,
		proxy:       "",
		sku:         "",
	}

	user.getProxy()
	user.client = client.GetTLS(user.proxy)
	user.loginAccount()

}

func (user *HibbettBase) getProxy() {
	file, err := os.Open("./configs/proxies.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Read the lines into a slice
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Select a random line
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(lines))
	user.proxy = lines[randomIndex]
	proxySplit := strings.Split(user.proxy, ":")

	user.proxy = "http://" + proxySplit[2] + ":" + proxySplit[3] + "@" + proxySplit[0] + ":" + proxySplit[1]

}

// func (user *HibbettBase) loginAccount() {

// 	type Credentials struct {
// 		Login    string `json:"login"`
// 		Password string `json:"password"`
// 	}

// 	creds := &Credentials{
// 		Login:    user.email,
// 		Password: user.password,
// 	}

// 	jsonData, err := json.Marshal(creds)
// 	if err != nil {
// 		fmt.Println("Error marshaling JSON: ", err)
// 		return
// 	}

// 	/*
// 		TLS Helper

// 		Args:
// 			client: TLS session
// 			method: HTTP method (GET, POST, etc)
// 			url: URL to request
// 			headers: HTTP headers
// 			body: HTTP body
// 			expectedResponse: 200, 201, etc
// 	*/

// 	// res := client.TlsRequest(client.TLSParams{
// 	// 	Client: user.client,
// 	// 	Method: http.MethodPost,
// 	// 	Url:    "https://hibbett-mobileapi.prolific.io/users/login",
// 	// 	Headers: http.Header{
// 	// 		"Accept":             {"*/*"},
// 	// 		"Accept-Encoding":    {"br;q=1.0, gzip;q=0.9, deflate;q=0.8"},
// 	// 		"Accept-Language":    {"en-US;q=1.0"},
// 	// 		"Connection":         {"keep-alive"},
// 	// 		"Content-Type":       {"application/json; charset=utf-8"},
// 	// 		"platform":           {"ios"},
// 	// 		"version":            {"6.3.0"},
// 	// 		"x-api-key":          {"0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0"},
// 	// 		"X-PX-AUTHORIZATION": {"4"}, //1 also works
// 	// 		"User-Agent":         {user.userAgent},
// 	// 	},
// 	// 	Body:             strings.NewReader(string(jsonData)),
// 	// 	ExpectedResponse: 200,
// 	// },
// 	// )

// 	if res != "error" {

// 		type Session struct {
// 			SessionID  string `json:"sessionID"`
// 			CustomerID string `json:"customerID"`
// 		}

// 		var responseData Session
// 		err = json.Unmarshal((res), &responseData)
// 		if err != nil {
// 			fmt.Println("Error parsing JSON:", err)
// 			Init(user.thread, user.account)
// 		} else {
// 			user.sessionId = responseData.SessionID
// 			user.customerId = responseData.CustomerID

// 			fmt.Println("Logged in")
// 			user.getPaymentId()
// 		}

// 	} else {
// 		fmt.Println("Error logging in")
// 		Init(user.thread, user.account)
// 	}

// }

// func (user *HibbettBase) getPaymentId() {

// 	res := client.TlsRequest(client.TLSParams{
// 		Client: user.client,
// 		Method: http.MethodGet,
// 		Url:    getPaymentIdUrlString(user.customerId),
// 		Headers: http.Header{
// 			"Accept":             {"*/*"},
// 			"Accept-Encoding":    {"br;q=1.0, gzip;q=0.9, deflate;q=0.8"},
// 			"Accept-Language":    {"en-US;q=1.0"},
// 			"Connection":         {"keep-alive"},
// 			"Content-Type":       {"application/json; charset=utf-8"},
// 			"platform":           {"ios"},
// 			"version":            {"6.3.0"},
// 			"Authorization":      {"Bearer " + user.sessionId},
// 			"x-api-key":          {"0PutYAUfHz8ozEeqTFlF014LMJji6Rsc8bpRBGB0"},
// 			"X-PX-AUTHORIZATION": {"4"}, //1 also works
// 			"User-Agent":         {user.userAgent},
// 		},
// 		Body:             nil,
// 		ExpectedResponse: 200,
// 	},
// 	)

// 	if res != "error" {

// 		// type Payment struct {
// 		// 	ID            string `json:"id"`
// 		// 	Type          string `json:"type"`
// 		// 	PaymentObject struct {
// 		// 		NameOnCard        string `json:"nameOnCard"`
// 		// 		CardType          string `json:"cardType"`
// 		// 		Number            string `json:"number"`
// 		// 		ExpirationMonth   int    `json:"expirationMonth"`
// 		// 		ExpirationYear    int    `json:"expirationYear"`
// 		// 		CreditCardExpired bool   `json:"creditCardExpired"`
// 		// 		CreditCardToken   string `json:"creditCardToken"`
// 		// 		EncryptedCVNValue string `json:"encryptedCVNValue"`
// 		// 	} `json:"paymentObject"`
// 		// }

// 		var payments []Payment
// 		err := json.Unmarshal([]byte(res), &payments)
// 		if err != nil {
// 			fmt.Println("Error parsing JSON:", err)
// 			Init(user.thread, user.account)
// 		} else {

// 			for _, payment := range payments {

// 				if payment.PaymentObject.Number == user.four {
// 					user.paymentId = payment.ID
// 					user.paymentType = payment.Type
// 				} else {
// 					fmt.Println("Error, no valid paymentID found")
// 				}
// 			}

// 		}

// 	} else {
// 		fmt.Println("Error getting payment id")
// 		Init(user.thread, user.account)
// 	}
// }
