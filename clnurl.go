package main

import (
	"github.com/elementsproject/glightning/glightning"
	"github.com/michael1011/clnurl/build"
	"github.com/michael1011/clnurl/clnurl"
	"log"
	"net/http"
	"os"
	"strconv"
)

var ln *glightning.Lightning
var plugin *glightning.Plugin
var cu *clnurl.ClnUrl

func main() {
	plugin = glightning.NewPlugin(onInit)
	ln = glightning.NewLightning()

	registerOptions(plugin)

	err := plugin.Start(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

func onInit(plugin *glightning.Plugin, _ map[string]glightning.Option, config *glightning.Config) {
	plugin.Log("Starting "+build.GetVersion(), glightning.Unusual)

	err := ln.StartUp(config.RpcFile, config.LightningDir)
	if err != nil {
		plugin.Log("Connection to node failed: "+err.Error(), glightning.Info)
		return
	}

	cfg := parseConfig(plugin)

	nodeBackend := &NodeBackend{lightning: ln}
	cu = clnurl.Init(cfg.cu, nodeBackend)

	addr := cfg.Host + ":" + strconv.Itoa(cfg.Port)
	plugin.Log("Starting HTTP server on: "+addr, glightning.Info)

	go func() {
		err := http.ListenAndServe(
			addr,
			registerRoutes(cfg),
		)
		if err != nil {
			plugin.Log("Starting HTTP server failed: "+err.Error(), glightning.Info)
		}
	}()
}
