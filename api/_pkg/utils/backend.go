package utils

import (
	"github.com/fiatjaf/makeinvoice"
)

type MakeInvoiceBackend struct {
	mkBackend makeinvoice.BackendParams
}

func (b *MakeInvoiceBackend) MakeInvoice(msats int64, description string) (string, error) {
	return makeinvoice.MakeInvoice(makeinvoice.Params{
		Msatoshi:           msats,
		Backend:            b.mkBackend,
		Description:        description,
		UseDescriptionHash: true,
	})
}
