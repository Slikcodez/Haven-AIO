package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"main/channels"
	"main/constants"
	"net/url"
	"strconv"
	"strings"
)

type Message struct {
	Size    float32 `json:"Size"`
	Sku     string  `json:"Sku"`
	Variant string  `json:"Varient"`
}

func getSizes() []float32 {
	content, err := ioutil.ReadFile("./configs/hibbett/sizes.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	// Split the content by line
	lines := strings.Split(string(content), "\n")
	// Create an empty slice to store the sizes
	var sizes []float32
	// Loop through the lines and convert each to a float32
	for _, line := range lines {
		line = strings.TrimSpace(line) // Remove leading/trailing whitespace
		if line == "" {
			continue // Skip empty lines
		}
		f, err := strconv.ParseFloat(line, 32)
		if err != nil {
			fmt.Println("Error converting to float32:", err)
			continue // Skip non-numeric lines
		}
		sizes = append(sizes, float32(f))
	}
	return sizes
}

func getSkus() []string {
	content, err := ioutil.ReadFile("./configs/hibbett/skus.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)

	}
	// Split the content by line
	lines := strings.Split(string(content), "\n")
	// Create an empty slice to store the sizes
	var skus []string
	// Loop through the lines and convert each to an integer
	for _, line := range lines {
		if line == "" {
			continue // Skip empty lines
		}
		skus = append(skus, line)
	}

	return skus
}

func ConnectHibbett() {

	for {
		skus := getSkus()

		u := url.URL{Scheme: "ws", Host: "38.102.8.15:12145", Path: ""}

		// Set up WebSocket connection
		c, _, _ := websocket.DefaultDialer.Dial(u.String(), nil)

		for {
			// Read incoming message
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read error:")
				break
			}

			// Extract JSON data from message
			var msg Message
			err = json.Unmarshal(message, &msg)
			if err != nil {
				log.Println("json error:", err)
				continue
			}
			log.Println(msg.Sku, "RESTOCKED")
			for _, valueSku := range skus {
				if strings.ToUpper(valueSku) == strings.ToUpper(msg.Sku) && constants.GlobalSettings.MinSize <= msg.Size {

					go func() {
						channels.HavenCloud <- fmt.Sprintf("%s:%f:%s", msg.Variant, msg.Size, msg.Sku)
					}()

				}
			}
		}
	}
}
