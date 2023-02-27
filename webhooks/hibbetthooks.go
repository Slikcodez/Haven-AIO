package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"main/constants"
	"net/http"
	"time"
)

type Embed struct {
	Title       string          `json:"title,omitempty"`
	Description string          `json:"description,omitempty"`
	URL         string          `json:"url,omitempty"`
	Timestamp   string          `json:"timestamp,omitempty"`
	Color       int             `json:"color,omitempty"`
	Footer      *EmbedFooter    `json:"footer,omitempty"`
	Image       *EmbedImage     `json:"image,omitempty"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`
	Author      *EmbedAuthor    `json:"author,omitempty"`
	Fields      []*EmbedField   `json:"fields,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

type EmbedFooter struct {
	Text    string `json:"text,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type EmbedImage struct {
	URL string `json:"url,omitempty"`
}

type EmbedThumbnail struct {
	URL string `json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name         string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

func SendWebhook(size string, sku string, price float64, number string, imageURL string, email string) error {
	// create a new Discord webhook message
	timestamp := time.Now().Format(time.RFC3339)

	// Create the message data as a struct
	message := struct {
		Content     interface{} `json:"content,omitempty"`
		Username    string      `json:"username,omitempty"`
		AvatarURL   string      `json:"avatar_url,omitempty"`
		Embeds      []*Embed    `json:"embeds,omitempty"`
		TTS         bool        `json:"tts,omitempty"`
		File        interface{} `json:"file,omitempty"`
		PayloadJSON interface{} `json:"payload_json,omitempty"`
	}{
		Embeds: []*Embed{
			{
				Title: "🎉 🥳 Haven Checkout Success 🥳 🎉",
				Color: 16777215,
				Fields: []*EmbedField{
					{
						Name:   "Sku",
						Value:  fmt.Sprintf("%s", sku),
						Inline: true,
					},
					{
						Name:   "Size",
						Value:  fmt.Sprintf("%s", size),
						Inline: true,
					},
					{
						Name:   "Price",
						Value:  fmt.Sprintf("%f", price),
						Inline: true,
					},
					{
						Name:   "Account",
						Value:  fmt.Sprintf("||%s||", email),
						Inline: true,
					},
					{
						Name:   "Order Number",
						Value:  fmt.Sprintf("||%s||", number),
						Inline: true,
					},
				},
				Footer: &EmbedFooter{
					Text: fmt.Sprintf("Haven Aio | Version %s", constants.Version),
				},
				Timestamp: timestamp,
				Thumbnail: &EmbedThumbnail{
					URL: imageURL,
				},
			},
		},
		Username:  "Haven",
		AvatarURL: "https://media.discordapp.net/attachments/935330170301186069/1079254579247583262/fotor_2023-2-25_22_11_41.png",
	}
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	webhookURL := constants.GlobalSettings.Webhook
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil

}
