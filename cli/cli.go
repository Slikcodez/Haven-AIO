package cli

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearCli() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
	return
}

func SelectRetailer() string {
	fmt.Print(`
██╗  ██╗ █████╗ ██╗   ██╗███████╗███╗   ██╗
██║  ██║██╔══██╗██║   ██║██╔════╝████╗  ██║
███████║███████║██║   ██║█████╗  ██╔██╗ ██║
██╔══██║██╔══██║╚██╗ ██╔╝██╔══╝  ██║╚██╗██║
██║  ██║██║  ██║ ╚████╔╝ ███████╗██║ ╚████║
╚═╝  ╚═╝╚═╝  ╚═╝  ╚═══╝  ╚══════╝╚═╝  ╚═══╝


Pick a retailer!
1. Hibbett
2. Coming Soon!

>>>  `)
	var module string
	_, err := fmt.Scanln(&module)
	if err != nil {
		return ""
	}

	//Temporary while only one module

	if module != "1" {
		fmt.Println("If you would like a new module, please suggest it!")
		ClearCli()
		return SelectRetailer()
	} else {
		return "1"
	}

}

func SelectTask() string {
	fmt.Print(`
Pick A Mode!
1. Hibbett Cloud
2. Hibbett Account Generator
3. Hibbett Initial

>>>  `)

	var mode string

	_, err := fmt.Scanln(&mode)
	if err != nil {
		return ""
	}

	if mode == "1" {
		return "1"
	} else {
		if mode == "3" {
			return "3"
		}
		return "2"
	}
}
