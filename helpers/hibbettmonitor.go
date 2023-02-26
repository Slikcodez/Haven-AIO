package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"main/channels"
	"net/url"
	"strings"
)

type Message struct {
	Size    string `json:"Size"`
	Sku     string `json:"Sku"`
	Variant string `json:"Varient"`
}

func getSizes() []string {
	content, err := ioutil.ReadFile("./configs/hibbett/sizes.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)

	}
	// Split the content by line
	lines := strings.Split(string(content), "\n")
	// Create an empty slice to store the sizes
	var sizes []string
	// Loop through the lines and convert each to an integer
	for _, line := range lines {
		if line == "" {
			continue // Skip empty lines
		}
		sizes = append(sizes, line)
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

	sizes := getSizes()
	skus := getSkus()

	u := url.URL{Scheme: "ws", Host: "38.102.8.15:8001", Path: ""}

	// Set up WebSocket connection
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:")
	}
	defer c.Close()
	fmt.Println("Connected to Haven Cloud Monitor")

	for {
		// Read incoming message
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		// Extract JSON data from message
		var msg Message
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("json error:", err)
			continue
		}

		for _, valueSku := range skus {
			if valueSku == msg.Sku {
				for _, valueSize := range sizes {
					if valueSize == msg.Size {
						go func() {
							channels.HavenCloud <- fmt.Sprintf("%s:%s:%s", msg.Variant, msg.Size, msg.Sku)
						}()
					}
				}
			}
		}

	}

}
