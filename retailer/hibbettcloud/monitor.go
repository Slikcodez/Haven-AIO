package hibbettcloud

import (
	"main/channels"
	"main/constants"
)

func (user *HibbettBase) Monitor() {
	constants.LogStatus(user.thread, "Listening For Restocks")

	channels.HavenCloud.On("restock", func(sku string) { user.preCart(sku) })

	for {
	}
}
