package bech32

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/btcsuite/btcutil/bech32"
	"github.com/michael1011/clnurl/api/_pkg/utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	cfg := utils.GetConfig()

	bits, err := bech32.ConvertBits([]byte(cfg.Endpoint+"/api/lnurlp"), 8, 5, true)
	if err != nil {
		utils.FormatError(w, http.StatusInternalServerError, err)
		return
	}

	lnurl, err := bech32.Encode("LNURL", bits)
	if err != nil {
		utils.FormatError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SetHeaders(w)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(strings.ToUpper(lnurl))
}
