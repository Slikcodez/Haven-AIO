package client

import (
	"io"

	http "github.com/bogdanfinn/fhttp"
)

type TLSParams struct {
	Client           HttpClient
	Method           string
	Url              string
	Headers          http.Header
	Body             io.Reader
	ExpectedResponse int
}
