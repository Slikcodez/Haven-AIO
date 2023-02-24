package webhook

import (
	"log"
	"testing"

	"github.com/gtuk/discordwebhook"
)

var itemLabel string = "Product"
var itemName string = "Jordan "
var sizeLabel string = "Size"
var sizeValue string = "9-11"
var storeLabel string = "Store"
var storeName string = "Hibbett"
var True bool = true
var WebhookSamples []WebhookStruct = []WebhookStruct{

	{
		Title:   "Successful Checkout from Nordstrom",
		Message: "Jordan 9-11's",
		Color:   "5294200",
		Webhook: "https://discord.com/api/webhooks/959531632082186290/U4_foTmEpjYPlPTafHtqlvLjv-_sDJH9k3-qdQdFh5tz0Enp4DMXx01PykwhfHFvM-0O",
		Fields: []discordwebhook.Field{
			{
				Name:  &sizeLabel,
				Value: &sizeValue,
			},
		},
	},

	{
		Title:   "Hibbett Haven",
		Color:   "5294200",
		Webhook: "https://discord.com/api/webhooks/959531632082186290/U4_foTmEpjYPlPTafHtqlvLjv-_sDJH9k3-qdQdFh5tz0Enp4DMXx01PykwhfHFvM-0O",
		Fields: []discordwebhook.Field{
			{
				Name:  &storeLabel,
				Value: &storeName,
			},

			{
				Name:   &itemLabel,
				Value:  &itemName,
				Inline: &True,
			},
			{
				Name:   &sizeLabel,
				Value:  &sizeValue,
				Inline: &True,
			},
		},
	},
}

func TestWebhook(t *testing.T) {

	for _, elem := range WebhookSamples {
		log.Println(elem.Fields)
		log.Println(elem.Webhook)
		DiscordWebhook(elem)
	}
}
