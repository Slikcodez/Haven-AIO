package webhook

import "github.com/gtuk/discordwebhook"

type WebhookStruct struct {
	Message string `json:"message"`
	Color   string `json:"color"`
	Title   string `json:"title"`
	// Key     string                 `json:"key"`
	Webhook string                 `json:"webhook"`
	Fields  []discordwebhook.Field `json:"fields"`
}

var botUsername string = "Test"
var authorName string = "Successful Checkout"
var botImage string = "https://media.discordapp.net/attachments/935389834279800843/1040468227223978044/Hibbett_Haven.png?width=1404&height=1404"

// var webhookString string = " "
