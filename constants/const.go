package constants

const ServerUrl string = `http://38.102.8.15`

const HibbettURL string = `https://hibbett-mobileapi.prolific.io/users/`

func GetPaymentIdUrlString(customerID string) string {
	return HibbettURL + customerID + "/payment_methods"
}
