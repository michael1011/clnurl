package main

import (
	"github.com/elementsproject/glightning/glightning"
	"github.com/michael1011/clnurl/clnurl"
	"github.com/michael1011/clnurl/clnurl/consts"
)

const (
	optionPrefix = "clnurl-"

	optionHost      = optionPrefix + "host"
	optionPort      = optionPrefix + "port"
	optionServeSite = optionPrefix + "serve-site"

	optionEndpoint           = optionPrefix + "endpoint"
	optionMinSendable        = optionPrefix + "min-sendable"
	optionMaxSendable        = optionPrefix + "max-sendable"
	optionInvoiceDescription = optionPrefix + "description"
)

type config struct {
	Host      string
	Port      int
	ServeSite bool

	cu *clnurl.Config
}

func parseConfig(p *glightning.Plugin) *config {
	host, _ := p.GetOption(optionHost)
	port, _ := p.GetIntOption(optionPort)
	serveSite, _ := p.GetBoolOption(optionServeSite)

	endpoint, _ := p.GetOption(optionEndpoint)
	description, _ := p.GetOption(optionInvoiceDescription)
	minSendable, _ := p.GetIntOption(optionMinSendable)
	maxSendable, _ := p.GetIntOption(optionMaxSendable)

	return &config{
		Host:      host,
		Port:      port,
		ServeSite: serveSite,

		cu: &clnurl.Config{
			Endpoint:           endpoint,
			InvoiceDescription: description,
			MinSendable:        int64(minSendable),
			MaxSendable:        int64(maxSendable),
		},
	}
}

func registerOptions(p *glightning.Plugin) {
	_ = p.RegisterNewOption(optionHost, "Host of the server", "127.0.0.1")
	_ = p.RegisterNewIntOption(optionPort, "Port of the server", 3000)
	_ = p.RegisterNewBoolOption(optionServeSite, "Whether the website should be served", true)

	_ = p.RegisterNewOption(optionEndpoint, "Publicly reachable endpoint of clnurl", consts.Endpoint)
	_ = p.RegisterNewIntOption(optionMinSendable, "Minimal sendable via the LNURL", consts.MinSendable)
	_ = p.RegisterNewIntOption(optionMaxSendable, "Maximal sendable via the LNURL", consts.MaxSendable)
	_ = p.RegisterNewOption(optionInvoiceDescription, "Description of the LNURL", consts.InvoiceDescription)
}
