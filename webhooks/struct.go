package webhook

type WebhookStruct struct {
	Message string `json:"message"`
	Color   string `json:"color"`
	Title   string `json:"title"`
	Key     string `json:"key"`
	Webhook string `json:"webhook"`
}
