package nordstrom

import (
	"io/ioutil"
	"net/http"
)

func getHeadersReq() (string, error) {
	url := "http://localhost:6942/api/getHeader"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
