package hibbettcloud

import (
	"bufio"
	"encoding/json"
	"fmt"
	http "github.com/bogdanfinn/fhttp"
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

	self := HibbettBase{
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

	self.getProxy()
	self.client = client.GetTLS(self.proxy)
	self.loginAccount()

}

func (self *HibbettBase) getProxy() {
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
	self.proxy = lines[randomIndex]
	proxySplit := strings.Split(self.proxy, ":")

	self.proxy = "http://" + proxySplit[2] + ":" + proxySplit[3] + "@" + proxySplit[0] + ":" + proxySplit[1]

}

func (self *HibbettBase) loginAccount() {

	type Credentials struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	creds := &Credentials{
		Login:    self.email,
		Password: self.password,
	}

	jsonData, err := json.Marshal(creds)
	if err != nil {
		fmt.Println("Error marshaling JSON: ", err)
		return
	}

	/*
		TLS Helper

		Args:
			client: TLS session
			method: HTTP method (GET, POST, etc)
			url: URL to request
			headers: HTTP headers
			body: HTTP body
			expectedResponse: 200, 201, etc
	*/

	res := client.TlsRequest(
		self.client,
		http.MethodPost,
		"https://hibbett-mobileapi.prolific.io/users/login",
		http.Header{
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
		strings.NewReader(string(jsonData)),
		200,
	)

	if res != "error" {

		type Session struct {
			SessionID  string `json:"sessionID"`
			CustomerID string `json:"customerID"`
		}

		var responseData Session
		err = json.Unmarshal([]byte(res), &responseData)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			Init(self.thread, self.account)
		} else {
			self.sessionId = responseData.SessionID
			self.customerId = responseData.CustomerID

			fmt.Println("Logged in")
			self.getPaymentId()
		}

	} else {
		fmt.Println("Error logging in")
		Init(self.thread, self.account)
	}

}

func (self *HibbettBase) getPaymentId() {

	res := client.TlsRequest(
		self.client,
		http.MethodGet,
		"https://hibbett-mobileapi.prolific.io/users/"+self.customerId+"/payment_methods",
		http.Header{
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
		nil,
		200,
	)

	if res != "error" {

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

		var payments []Payment
		err := json.Unmarshal([]byte(res), &payments)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			Init(self.thread, self.account)
		} else {

			for _, payment := range payments {

				if payment.PaymentObject.Number == self.four {
					self.paymentId = payment.ID
					self.paymentType = payment.Type
				} else {
					fmt.Println("Error, no valid paymentID found")
				}
			}

		}

	} else {
		fmt.Println("Error getting payment id")
		Init(self.thread, self.account)
	}
}
