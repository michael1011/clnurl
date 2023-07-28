package main

import (
	"github.com/elementsproject/glightning/glightning"
	"math/rand"
	"strconv"
	"time"
)

type NodeBackend struct {
	lightning *glightning.Lightning
}

func (n *NodeBackend) Disconnect() error {
	return nil
}

func (n *NodeBackend) MakeInvoice(msats int64, description string) (string, error) {
	inv, err := n.lightning.Invoice(uint64(msats), makeRandomLabel(), description)
	if err != nil {
		return "", err
	}

	return inv.Bolt11, err
}

func makeRandomLabel() string {
	return "makeinvoice/" + strconv.FormatInt(time.Now().Unix()+rand.Int63(), 16)
}
