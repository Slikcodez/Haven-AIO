package constants

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const ServersUrl string = `http://38.102.8.15`
const HibbettURL string = `https://hibbett-mobileapi.prolific.io/users/`

const BearerToken string = `pk_cqzp7m2w4zsl5jlx233swz3uj7prgkmc5lmzg205`

func LogStatus(thread string, message string) {
	fmt.Println("Thread " + thread + ": " + message)
}

func GetPaymentIdUrlString(customerID string) string {
	return HibbettURL + customerID + "/payment_methods"
}

func UnmarshalRequestError(req string, resptype string) (string, error) {
	type response struct {
		StatusCode int    `json:"statusCode"`
		Body       []byte `json:"body"`
	}

	var resp response
	err := json.Unmarshal([]byte(req), &resp)
	if err != nil {
		return "", err
	}

	if resptype == "body" {
		return string(resp.Body), nil
	} else {
		return strconv.Itoa(resp.StatusCode), nil
	}
}
