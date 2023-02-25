package hibbettcloud

import (
	"main/constants"
)

func (user *HibbettBase) Monitor() {
	constants.LogStatus(user.thread, "Listening For Restocks")
	for {
		sku := <-HavenCloud

		go func() {
			user.preCart(sku)
		}()
	}
}
