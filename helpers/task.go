package helpers

import (
	"bufio"
	"fmt"
	"main/retailer/hibbettcloud"
	"os"
	"strconv"
	"sync"
	"time"
)

func readAccountsFile() ([]string, error) {
	file, err := os.Open("./configs/hibbett/accounts.txt")
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Crash: Please check your accounts.txt file.")
			os.Exit(1)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var accounts []string
	for scanner.Scan() {
		accounts = append(accounts, scanner.Text())
	}

	return accounts, nil
}

func InitTask() {

	go func() {
		ConnectHibbett()
	}()
	time.Sleep(500 * time.Millisecond)

	var accounts, _ = readAccountsFile()

	var wg sync.WaitGroup
	ThreadNumTask := 1
	for _, account := range accounts {
		wg.Add(1)
		go func() {
			defer wg.Done()
			hibbettcloud.Init(strconv.Itoa(ThreadNumTask), account, "")

		}()
		time.Sleep(500 * time.Millisecond)
		ThreadNumTask++
	}
	wg.Wait()
}

func InitInitTask() {

	var accounts, _ = readAccountsFile()

	var wg sync.WaitGroup
	ThreadNumTask := 1
	for _, account := range accounts {
		wg.Add(1)
		go func() {
			defer wg.Done()
			hibbettcloud.Init(strconv.Itoa(ThreadNumTask), account, "initial")

		}()
		time.Sleep(500 * time.Millisecond)
		ThreadNumTask++
	}
	wg.Wait()
}
