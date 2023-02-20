package hibbettcloud

import "encoding/json"

type HibbettPayments struct {
	Id            string        `json:"id"`
	Type          string        `json:"type"`
	PaymentObject PaymentObject `json:"paymentObject"`
}

type PaymentObject struct {
	Number string `json:"number"`
}

func (self *HibbettBase) getHibbetPaymentOptions(response string) (data []HibbettPayments) {

	json.Unmarshal([]byte(response), &data)
	return
}
