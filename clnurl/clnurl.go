package clnurl

import (
	"fmt"

	"github.com/fiatjaf/makeinvoice"
)

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
		Msatoshi:           msats,
		Backend:            cu.backend,
		Description:        cu.getMetaData(),
		UseDescriptionHash: true,
	})
	if err != nil {
		return nil, err
	}

	return &InvoiceResponse{
		Invoice: invoice,
		Routes:  []string{},
	}, nil
}
