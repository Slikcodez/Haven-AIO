package webhook

import (
	"log"

	"github.com/gtuk/discordwebhook"
)

func DiscordWebhook(wh WebhookStruct) {

	embeds := discordwebhook.Embed{
		Title:       &wh.Title,
		Description: &wh.Message,
		Color:       &wh.Color,
		Fields:      &wh.Fields,
	}
	// log.Println(embeds)
	embedArray := []discordwebhook.Embed{embeds}
	log.Println(embedArray)
	message := discordwebhook.Message{
		Username:  &botUsername,
		AvatarUrl: &botImage,
		Embeds:    &embedArray,
	}

	err := discordwebhook.SendMessage(webhookString, message)
	if err != nil {
		log.Println(err)
		log.Fatalf("Failure in webhook")

	}
}
