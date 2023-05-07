package auth

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"main/constants"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

func authLoop(key string) {
	for {
		reqBody, err := json.Marshal(map[string]string{
			"license": key,
			"hwid":    hex.EncodeToString(uuid.NodeID()),
			"type":    "check",
		})
		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s:4567/auth", constants.ServersUrl), bytes.NewBuffer(reqBody))
		if err != nil {
			panic(err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", constants.BearerToken))

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			continue
		}

		if res.StatusCode == http.StatusOK {
			time.Sleep(120 * time.Second)
			continue
		} else {
			os.Exit(0)
		}
	}
}

func StartAuthFunc(key string) (status string) {
	for {
		reqBody, err := json.Marshal(map[string]string{
			"license": key,
			"hwid":    hex.EncodeToString(uuid.NodeID()),
			"type":    "init",
		})
		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s:4567/auth", constants.ServersUrl), bytes.NewBuffer(reqBody))
		if err != nil {
			panic(err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", constants.BearerToken))

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			continue
		}

		if res.StatusCode == http.StatusOK {
			go func() {
				authLoop(key)
			}()
			return "200"
		} else {
			return "403"
		}
	}
}
