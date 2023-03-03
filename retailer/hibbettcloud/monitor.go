package hibbettcloud

import (
	"main/channels"
	"main/constants"
)

func (user *HibbettBase) Monitor() {
	constants.LogStatus(user.thread, "Listening For Restocks")

	for {
		sku := <-channels.HavenCloud
		user.preCart(sku)
	}
}
