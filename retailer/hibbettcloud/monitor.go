package hibbettcloud

import (
	"main/channels"
	"main/constants"
)

func (user *HibbettBase) recieve(sku string) {
	user.preCart(sku)
}

func (user *HibbettBase) Monitor() {
	constants.LogStatus(user.thread, "Listening For Restocks")
	for {
		event := <-channels.HavenCloud.Once(constants.RandString())
		user.preCart(event.Args[0].(string))
	}

}
