package main

import (
	"fmt"
	"log"
	"main/auth"
	"main/cli"
	"main/constants"
	"main/helpers"
	"os"
	"time"
)

func LaunchCli() {

	var res string
	res = cli.SelectRetailer()

	if res == "1" {

		res = cli.SelectTask()

		if res == "1" {
			helpers.InitTask()
		} else {
			fmt.Println("Not Incorporated yet")
		}

	} else {
		fmt.Println("If you would like a new module, please suggest it!")
		cli.ClearCli()
		LaunchCli()
	}

}

func main() {

	// Defines webhook and key values in constants.go file
	settings, err := constants.ReadSettingsFile("./configs/settings.json")
	if err != nil {
		log.Fatal(err)
	}

	if auth.StartAuthFunc(settings.Key) != "200" {
		fmt.Println("Error Authenticating Haven")
		time.Sleep(5 * time.Second)
		os.Exit(0)

	}
	LaunchCli()

}
