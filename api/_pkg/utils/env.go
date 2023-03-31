package utils

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/fiatjaf/makeinvoice"
	"github.com/michael1011/clnurl/clnurl"
	"github.com/michael1011/clnurl/clnurl/consts"
)

const (
	envEndpoint           = "ENDPOINT"
	envInvoiceDescription = "INVOICE_DESCRIPTION"

	envMinSendable = "MIN_SENDABLE"
	envMaxSendable = "MAX_SENDABLE"

	envClnNodeId = "CLN_NODE_ID"
	envClnRune   = "CLN_RUNE"
	envClnHost   = "CLN_HOST"
)

var ErrMissingEnv = errors.New("env variable missing")

func GetConfig() *clnurl.Config {
	parseInt64 := func(val string) (int64, error) {
		return strconv.ParseInt(val, 10, 64)
	}

	isEmptyInt64 := func(val int64) bool {
		return val == 0
	}

	parseString := func(val string) (string, error) { return val, nil }
	isEmptyString := func(val string) bool { return val == "" }

	return &clnurl.Config{
		MinSendable: getEnvVarOptional[int64](
			envMinSendable,
			consts.MinSendable,
			parseInt64,
			isEmptyInt64,
		),
		MaxSendable: getEnvVarOptional[int64](
			envMaxSendable,
			consts.MaxSendable,
			parseInt64,
			isEmptyInt64,
		),
		Endpoint: strings.TrimSuffix(getEnvVarOptional[string](
			envEndpoint,
			consts.Endpoint,
			parseString,
			isEmptyString,
		), "/"),
		InvoiceDescription: getEnvVarOptional[string](
			envInvoiceDescription,
			consts.InvoiceDescription,
			parseString,
			isEmptyString,
		),
	}
}

func GetClnBackend() (*makeinvoice.CommandoParams, error) {
	nodeId, err := getEnvVar(envClnNodeId)
	if err != nil {
		return nil, err
	}

	clnRune, err := getEnvVar(envClnRune)
	if err != nil {
		return nil, err
	}

	host, err := getEnvVar(envClnHost)
	if err != nil {
		return nil, err
	}

	return &makeinvoice.CommandoParams{
		NodeId: nodeId,
		Rune:   clnRune,
		Host:   host,
	}, nil
}

func getEnvVar(key string) (string, error) {
	val := os.Getenv(key)

	if val == "" {
		return "", ErrMissingEnv
	}

	return val, nil
}

func getEnvVarOptional[T any](
	key string,
	defaultVal T,
	parse func(val string) (T, error),
	isEmpty func(val T) bool,
) T {
	val, err := getEnvVar(key)
	if err != nil {
		return defaultVal
	}

	parsed, err := parse(val)
	if err != nil {
		return defaultVal
	}

	if isEmpty(parsed) {
		return defaultVal
	}

	return parsed
}
