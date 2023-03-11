package hibbettcloud

import (
	"bufio"
	"fmt"
	"main/client"
	"main/constants"
	"math/rand"
	"os"
	"strings"
	"time"
)

type HibbettBase struct {
	client           client.HttpClient
	thread           string
	account          string
	paymentId        string
	token            string
	userAgent        string
	paymentType      string
	sessionId        string
	sessionEX        string
	customerId       string
	email            string
	password         string
	four             string
	checkedOut       bool
	genningPx        bool
	proxy            string
	initialSku       string
	detected         bool
	sku              string
	cartId           string
	cvv              string
	mode             string
	headerPlaceOrder string
}

func Init(thread string, account string, mode string) {

	user := HibbettBase{
		thread:    thread,
		account:   account,
		mode:      mode,
		paymentId: "",
		token:     "",
		//userAgent:   fmt.Sprintf("Hibbett | CG/6.3.0 (com.hibbett.hibbett-sports; build:%v; iOS 16.0.3) Alamofire/5.0.0-rc.3", rand.Intn(15000)+1),
		userAgent: fmt.Sprintf("Hibbett" + constants.RandString() + constants.RandString()),

		paymentType:      "",
		sessionId:        "",
		sessionEX:        "",
		headerPlaceOrder: "",
		customerId:       "",
		email:            strings.Split(account, ":")[0],
		password:         strings.Split(account, ":")[1],
		four:             strings.Split(account, ":")[2],
		cvv:              strings.Split(account, ":")[3],
		checkedOut:       false,
		proxy:            "",
		sku:              "",
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
