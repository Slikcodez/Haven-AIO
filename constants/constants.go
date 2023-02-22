package constants

import "fmt"

const ServersUrl string = `http://38.102.8.15`
const HibbettURL string = `https://hibbett-mobileapi.prolific.io/users/`

const BearerToken string = `pk_cqzp7m2w4zsl5jlx233swz3uj7prgkmc5lmzg205`

func LogStatus(thread string, message string) {
	fmt.Println("Thread " + thread + ": " + message)
}

func GetPaymentIdUrlString(customerID string) string {
	return HibbettURL + customerID + "/payment_methods"
}
