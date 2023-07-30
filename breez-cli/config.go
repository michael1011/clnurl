package main

import (
	"fmt"
	"github.com/michael1011/clnurl/clnurl/consts"
	"github.com/spf13/viper"
	"os"
)

type config struct {
	Address  string
	Endpoint string
	Mnemonic string
	ApiKey   string

	InvoiceDescription string

	MinSendable int64
	MaxSendable int64
}

var cfg config

func parseConfig() {
	viper.SetDefault("InvoiceDescription", "Breez SDK LNURL")
	viper.SetDefault("MinSendable", consts.MinSendable)
	viper.SetDefault("MaxSendable", consts.MaxSendable)

	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Could not read config file: " + err.Error())
		os.Exit(1)
		return
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Println("Could not decode config into struct: " + err.Error())
		os.Exit(1)
		return
	}

	if addressFlag != "" {
		cfg.Address = addressFlag
	}

	if cfg.Endpoint == "" {
		cfg.Endpoint = "http://" + cfg.Address
	}
}
