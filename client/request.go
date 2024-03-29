package client

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/url"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
)

type HttpClient interface {
	GetCookies(u *url.URL) []*http.Cookie
	SetCookies(u *url.URL, cookies []*http.Cookie)
	SetCookieJar(jar http.CookieJar)
	SetProxy(proxyUrl string) error
	GetProxy() string
	SetFollowRedirect(followRedirect bool)
	GetFollowRedirect() bool
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (resp *http.Response, err error)
	Head(url string) (resp *http.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

func GetTLS(proxy string) HttpClient {
	jar := tls_client.NewCookieJar()

	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(tls_client.Okhttp4Android7),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
	}

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = client.SetProxy(proxy)
	if err != nil {
		return nil
	}

	return client
}

func TlsRequest(params TLSParams) ([]byte, error) {

	req, err := http.NewRequest(params.Method, params.Url, params.Body)
	if err != nil {
		return nil, err
	}

	req.Header = params.Headers

	resp, err := params.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		Body.Close()
	}(resp.Body)

	readBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if params.ExpectedResponse == resp.StatusCode {
		return readBytes, nil
	} else {
		type response struct {
			StatusCode int    `json:"statusCode"`
			Body       []byte `json:"body"`
		}
		errorMsg := response{StatusCode: resp.StatusCode, Body: readBytes}
		jsonBytes, err := json.Marshal(errorMsg)
		if err != nil {
			return nil, err
		}
		errorMsgString := string(jsonBytes)
		return nil, errors.New(errorMsgString)
	}

}
