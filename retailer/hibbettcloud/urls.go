package hibbettcloud

const serverUrl string = `https://hibbett-mobileapi.prolific.io/users/`


func getPaymentIdUrlString(customerID string) string {
	return serverUrl + customerID + "/payment_methods"
}
