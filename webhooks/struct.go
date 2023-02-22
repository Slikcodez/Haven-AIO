package webhook

import "github.com/gtuk/discordwebhook"

type WebhookStruct struct {
	Message string                 `json:"message"`
	Color   string                 `json:"color"`
	Title   string                 `json:"title"`
	Key     string                 `json:"key"`
	Webhook string                 `json:"webhook"`
	Fields  []discordwebhook.Field `json:"fields"`
}

var botUsername string = "Test"
var botImage string = "https://upload.wikimedia.org/wikipedia/en/thumb/e/e8/Shell_logo.svg/2560px-Shell_logo.svg.png"
var webhookString string = " "
