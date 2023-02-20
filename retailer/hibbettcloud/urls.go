package hibbettcloud

const ServerURL string = `https://hibbett-mobileapi.prolific.io/users/`

func getPaymentIdUrlString(customerID string) string {
	return ServerURL + customerID + "/payment_methods"
}
