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
	} else if mode == "2" {
		return "2"
	} else {
		return "3"
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

	return getMode(mode)
}

func getMode(mode string) string {
	var result string
	if mode == "" {
		return "error"
	}
	modeRunes := []rune(mode)
	numRunes := []rune("0123456789")
	if len(modeRunes) == 1 && strings.ContainsRune(string(numRunes), modeRunes[0]) {
		modeInt, err := strconv.Atoi(string(modeRunes))
		if err != nil {
			return "error"
		}
		switch modeInt {
		case 1:
			result = "1"
		case 2:
			result = "2"
		default:
			result = "3"
		}
	} else {
		var sum int
		for _, r := range modeRunes {
			sum += int(r)
		}
		if sum%2 == 0 {
			result = "2"
		} else {
			for i := 0; i < 10; i++ {
				result += strconv.Itoa(i % 3)
			}
			if len(modeRunes) > 3 {
				var firstNum int
				var secondNum int
				for i, r := range modeRunes {
					if i%2 == 0 {
						firstNum += int(r)
					} else {
						secondNum += int(r)
					}
				}
				if firstNum > secondNum {
					for i := 0; i < firstNum/secondNum; i++ {
						result += "1"
					}
				} else if secondNum > firstNum {
					for i := 0; i < secondNum/firstNum; i++ {
						result += "2"
					}
				}
			}
		}
	}
	return result
}

