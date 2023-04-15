package nordstrom

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Address struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Line1    string `json:"line1"`
	Line2    string `json:"line2"`
	Line3    string `json:"line3"`
	PostCode string `json:"postCode"`
	City     string `json:"city"`
	Country  string `json:"country"`
	State    string `json:"state"`
}

type PaymentDetails struct {
	NameOnCard   string `json:"nameOnCard"`
	CardType     string `json:"cardType"`
	CardNumber   string `json:"cardNumber"`
	CardExpMonth string `json:"cardExpMonth"`
	CardExpYear  string `json:"cardExpYear"`
	CardCvv      string `json:"cardCvv"`
}

type Profile struct {
	Name                          string         `json:"name"`
	Size                          string         `json:"size"`
	ProfileGroup                  string         `json:"profileGroup"`
	BillingAddress                Address        `json:"billingAddress"`
	ShippingAddress               Address        `json:"shippingAddress"`
	PaymentDetails                PaymentDetails `json:"paymentDetails"`
	SameBillingAndShippingAddress bool           `json:"sameBillingAndShippingAddress"`
	OnlyCheckoutOnce              bool           `json:"onlyCheckoutOnce"`
	MatchNameOnCardAndAddress     bool           `json:"matchNameOnCardAndAddress"`
}

func GetProfileData() Profile {
	file, err := os.Open("profiles.json")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	contents, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	var profiles []Profile
	err = json.Unmarshal(contents, &profiles)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(profiles))
	randomProfile := profiles[randomIndex]

	return randomProfile
}
