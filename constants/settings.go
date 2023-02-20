package constants

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Settings struct {
	Key     string `json:"key"`
	Webhook string `json:"webhook"`
}

// ReadSettingsFile reads a JSON file located at the given filepath and
// returns a Settings object with the values from the file.
func ReadSettingsFile(filepath string) (Settings, error) {
	// Open the JSON file
	file, err := os.Open(filepath)
	if err != nil {
		return Settings{}, err
	}
	defer file.Close()

	// Read the contents of the file
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return Settings{}, err
	}

	// Unmarshal the JSON data into a Settings object
	var settings Settings
	if err := json.Unmarshal(contents, &settings); err != nil {
		return Settings{}, err
	}

	return settings, nil
}
