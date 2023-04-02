package clnurl

import (
	"fmt"
)

func Init(cfg *Config, backend Backend) *ClnUrl {
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

	invoice, err := cu.backend.MakeInvoice(msats, cu.getMetaData())
	if err != nil {
		return nil, err
	}

	return &InvoiceResponse{
		Invoice: invoice,
		Routes:  []string{},
	}, nil
}
