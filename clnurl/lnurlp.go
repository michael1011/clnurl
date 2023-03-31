package clnurl

import (
	"encoding/json"
	"fmt"
)

const (
	metadataTextPlain = "text/plain"
)

func (cu *ClnUrl) getMetaData() string {
	metdata := [][2]string{
		{
			metadataTextPlain,
			cu.cfg.InvoiceDescription,
		},
	}

	meta, _ := json.Marshal(metdata)
	return string(meta)
}

func (cu *ClnUrl) GetLnurlp() (Lnurlp, error) {
	return Lnurlp{
		Tag:         tagPayRequest,
		MaxSendable: cu.cfg.MaxSendable,
		MinSendable: cu.cfg.MinSendable,
		Metadata:    cu.getMetaData(),
		Callback:    fmt.Sprintf("%s/api/invoice", cu.cfg.Endpoint),
	}, nil
}
