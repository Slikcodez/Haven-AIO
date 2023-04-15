package hibbettcloud

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	"main/constants"
	"math/rand"
	"time"
)

type Proxy struct {
	Proxy string `json:"proxy"`
}

type Response struct {
	Error     bool   `json:"error"`
	UserAgent string `json:"userAgent"`
	Px3       string `json:"_px3"`
}

func (user *HibbettBase) GetPX(productData string) {
	url := "http://38.102.8.15:8089/genPX" // Replace with your desired URL
	proxy := &Proxy{
		Proxy: user.proxy,
	}
	jsonData, err := json.Marshal(proxy)
	if err != nil {
		fmt.Println("Error marshaling proxy:", err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error getting px:", err)
		user.loginAccount()
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error getting px:", err)
		user.loginAccount()
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("%d\n while attempting to get px", resp.StatusCode)
	} else {
		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			fmt.Println("Error decoding response:", err)
			user.loginAccount()
		}

		if response.Error {
			user.pxToken = generateRandomString(100, 150)
		} else {
			user.userAgent = response.UserAgent
			user.pxToken = response.Px3
		}

		constants.LogStatus(user.thread, "Baked A Cookie")
		user.preCart(productData)
	}
}

func generateRandomString(min, max int) string {
	length := min + rand.Intn(max-min)
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(b)[:length]
}
