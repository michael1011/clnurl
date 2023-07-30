package main

import (
	"fmt"
	"github.com/michael1011/clnurl/breez"
	"github.com/michael1011/clnurl/build"
	"github.com/michael1011/clnurl/clnurl"
	"github.com/michael1011/clnurl/router"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var (
	addressFlag string
	configFile  string
)

func main() {
	var br *breez.Backend

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "starts the LNURL server",
		PreRun: func(_ *cobra.Command, _ []string) {
			parseConfig()
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting " + build.GetVersion())

			var err error
			br, err = breez.Init(
				"",
				cfg.Mnemonic,
				cfg.ApiKey,
				true,
			)

			if err != nil {
				fmt.Println("Breez init failed: " + err.Error())
				os.Exit(1)
				return
			}

			info, err := br.NodeInfo()
			if err != nil {
				fmt.Println("Could not get node info: " + err.Error())
				os.Exit(1)
				return
			}

			fmt.Println("Connected to Breez: " + info.Id)

			c := make(chan os.Signal)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)
			go func() {
				<-c

				fmt.Println("Shutting down")
				cleanup(br)

				os.Exit(0)
			}()

			defer func() {
				cleanup(br)
			}()

			cu := clnurl.Init(&clnurl.Config{
				Endpoint:           cfg.Endpoint,
				InvoiceDescription: cfg.InvoiceDescription,
				MinSendable:        cfg.MinSendable,
				MaxSendable:        cfg.MaxSendable,
			}, br)

			fmt.Println("Starting HTTP server on: http://" + cfg.Address)

			err = router.Start(cu, cfg.Address, true)
			if err != nil {
				fmt.Println("Starting HTTP server failed: " + err.Error())
			}
		},
	}

	infoCmd := &cobra.Command{
		Use:   "info",
		Short: "shows info about the Breez node",
		PreRun: func(_ *cobra.Command, _ []string) {
			parseConfig()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			br, err = breez.Init(
				"",
				cfg.Mnemonic,
				cfg.ApiKey,
				true,
			)
			if err != nil {
				return err
			}

			info, err := br.NodeInfo()
			if err != nil {
				return err
			}

			fmt.Println("ID:      " + info.Id)
			fmt.Println("Balance: " + strconv.FormatFloat(float64(info.ChannelsBalanceMsat)/1000, 'f', -1, 64) + " sats")

			return nil
		},
	}

	rootCmd := &cobra.Command{
		Use:     "breez",
		Version: build.GetVersion(),
		PersistentPostRun: func(_ *cobra.Command, _ []string) {
			cleanup(br)
		},
	}

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "breez.toml", "path to the config file")
	rootCmd.PersistentFlags().StringVarP(&addressFlag, "address", "a", "", "host:port to bind to")
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(infoCmd)
	err := rootCmd.Execute()

	if err != nil {
		fmt.Println("Could not start application: " + err.Error())
		os.Exit(1)
	}
	return

}

func cleanup(br *breez.Backend) {
	if br == nil {
		return
	}

	err := br.Terminate()
	if err != nil {
		fmt.Println("Disconnect failed: " + err.Error())
		os.Exit(1)
	}
}
