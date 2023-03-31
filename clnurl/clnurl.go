package clnurl

import (
	"fmt"

	"github.com/fiatjaf/makeinvoice"
)

type Config struct {
	MinSendable int64
	MaxSendable int64
}

type ClnUrl struct {
	cfg     *Config
	backend makeinvoice.BackendParams
}

type InvoiceResponse struct {
	Invoice string   `json:"pr"`
	Routes  []string `json:"routes"`
}

func Init(cfg *Config, backend makeinvoice.BackendParams) *ClnUrl {
	return &ClnUrl{
		cfg:     cfg,
		backend: backend,
	}
}

func (cu *ClnUrl) MakeInvoice(msats int64) (*InvoiceResponse, error) {
	if msats < cu.cfg.MinSendable || msats > cu.cfg.MaxSendable {
		return nil, fmt.Errorf(
			"%d not in bounds of %d - %d",
			msats,
			cu.cfg.MinSendable,
			cu.cfg.MaxSendable,
		)
	}

	invoice, err := makeinvoice.MakeInvoice(makeinvoice.Params{
		Msatoshi: msats,
		Backend:  cu.backend,
	})
	if err != nil {
		return nil, err
	}

	return &InvoiceResponse{
		Invoice: invoice,
		Routes:  []string{},
	}, nil
}
