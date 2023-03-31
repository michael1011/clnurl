package utils

import (
	"errors"
	"os"

	"github.com/fiatjaf/makeinvoice"
)

const (
	clnNodeId = "CLN_NODE_ID"
	clnRune   = "CLN_RUNE"
	clnHost   = "CLN_HOST"
)

var ErrMissingEnv = errors.New("env variable missing")

func GetClnBackend() (*makeinvoice.CommandoParams, error) {
	nodeId, err := getEnvVar(clnNodeId)
	if err != nil {
		return nil, err
	}

	nodeRune, err := getEnvVar(clnRune)
	if err != nil {
		return nil, err
	}

	host, err := getEnvVar(clnHost)
	if err != nil {
		return nil, err
	}

	return &makeinvoice.CommandoParams{
		NodeId: nodeId,
		Rune:   nodeRune,
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
