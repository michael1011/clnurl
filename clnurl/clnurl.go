package clnurl

import (
	"github.com/fiatjaf/makeinvoice"
)

type ClnUrl struct {
	backend makeinvoice.BackendParams
}

func Init(backend makeinvoice.BackendParams) *ClnUrl {
	return &ClnUrl{
		backend: backend,
	}
}

func (cu *ClnUrl) MakeInvoice(msats int64, description string) (string, error) {
	return makeinvoice.MakeInvoice(makeinvoice.Params{
		Msatoshi:    msats,
		Description: description,
		Backend:     cu.backend,
	})
}
