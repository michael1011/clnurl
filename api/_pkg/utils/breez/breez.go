package breez

import (
	breez "github.com/breez/breez-sdk-go/breez_sdk"
)

type Backend struct {
	db *dbSync

	sdk *breez.BlockingBreezServices
}

type breezListener struct{}

func (breezListener) Log(breez.LogEntry) {}

func (breezListener) OnEvent(breez.BreezEvent) {}

func Init(postgresUrl, breezMnemonic, breezApiKey string) (*Backend, error) {
	db, err := initDbSync(postgresUrl)
	if err != nil {
		return nil, err
	}

	if err = db.download(); err != nil {
		return nil, err
	}

	bl := breezListener{}
	_ = breez.SetLogStream(bl)

	seed, err := breez.MnemonicToSeed(breezMnemonic)
	if err != nil {
		return nil, err
	}

	config := breez.DefaultConfig(breez.EnvironmentTypeProduction, breezApiKey, breez.NodeConfigGreenlight{})
	sdk, err := breez.Connect(config, seed, bl)
	if err != nil {
		return nil, err
	}

	return &Backend{
		db:  db,
		sdk: sdk,
	}, nil
}

func (b *Backend) Disconnect() error {
	if err := b.db.upload(); err != nil {
		return err
	}

	if err := b.db.close(); err != nil {
		return err
	}

	return b.sdk.Disconnect()
}

func (b *Backend) MakeInvoice(msats int64, description string) (string, error) {
	inv, err := b.sdk.ReceivePayment(uint64(msats), description)
	if err != nil {
		return "", err
	}

	return inv.Bolt11, nil
}
