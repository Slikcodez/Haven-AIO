package webhook

import (
	"log"
	"main/constants"

	"github.com/gtuk/discordwebhook"
)

func DiscordWebhook(wh WebhookStruct) {

	embeds := discordwebhook.Embed{
		Description: &wh.Message,
		Color:       &wh.Color,
		Fields:      &wh.Fields,
		Author: &discordwebhook.Author{
			Name:    &authorName,
			IconUrl: &botImage,
		},
	}
	embedArray := []discordwebhook.Embed{embeds}
	message := discordwebhook.Message{
		Username:  &botUsername,
		AvatarUrl: &botImage,
		Embeds:    &embedArray,
	}

	if (wh.Webhook) == "" {
		err := discordwebhook.SendMessage(constants.GlobalSettings.Webhook, message)
		if err != nil {
			log.Println(err)
		}
	} else {
		err := discordwebhook.SendMessage(wh.Webhook, message)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("successful webhook")
		}
	}
}
