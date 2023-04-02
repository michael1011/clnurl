package clnurl

import (
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/bech32"
	"strings"
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

func (cu *ClnUrl) GetLnurlpBech32() (string, error) {
	bits, err := bech32.ConvertBits([]byte(cu.cfg.Endpoint+"/api/lnurlp"), 8, 5, true)
	if err != nil {
		return "", err
	}

	lnurl, err := bech32.Encode("LNURL", bits)
	if err != nil {
		return "", err
	}

	return strings.ToUpper(lnurl), nil
}
