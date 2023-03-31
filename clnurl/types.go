package clnurl

import "github.com/fiatjaf/makeinvoice"

const (
	tagPayRequest LnurlTag = "payRequest"
)

type LnurlTag string

type Config struct {
	Endpoint           string
	InvoiceDescription string
	MinSendable        int64
	MaxSendable        int64
}

type ClnUrl struct {
	cfg     *Config
	backend makeinvoice.BackendParams
}

type InvoiceResponse struct {
	Invoice string   `json:"pr"`
	Routes  []string `json:"routes"`
}

type Lnurlp struct {
	Callback    string   `json:"callback"`
	MaxSendable int64    `json:"maxSendable"`
	MinSendable int64    `json:"minSendable"`
	Metadata    string   `json:"metadata"`
	Tag         LnurlTag `json:"tag"`
}
