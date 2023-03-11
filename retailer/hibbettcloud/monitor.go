package hibbettcloud

import (
	"bufio"
	"main/channels"
	"main/constants"
	"math/rand"
	"os"
	"time"
)

func GetRandVar() (string, error) {
	// Open the file
	file, err := os.Open("./configs/hibbett/initial.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the lines of the file into a slice
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Select a random line from the slice
	rand.Seed(time.Now().Unix())
	randomLine := lines[rand.Intn(len(lines))]

	return randomLine, nil
}

func (user *HibbettBase) Monitor() {

	if user.mode != "" {

		for {
			if time.Now().Hour() == constants.GlobalSettings.StartTime {
				varient, err := GetRandVar()
				if err == nil {
					user.preCart(varient)
				}

			}
		}
	} else {

		constants.LogStatus(user.thread, "Listening For Restocks")
		for {
			event := <-channels.HavenCloud.Once(constants.RandString())
			user.preCart(event.Args[0].(string))
		}
	}
}
