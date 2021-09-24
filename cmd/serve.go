package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagBind        = "bind"
	flagBindShort   = "b"
	flagBindDefault = "0.0.0.0"
	flagBindDescr   = "bind address for web server"

	flagPort        = "port"
	flagPortShort   = "p"
	flagPortDefault = 1323
	flagPortDescr   = "port for web server"

	envFlagBind       = flagBind
	envFlagPort       = flagPort
)

var serveHTTPCommand = &cobra.Command{
	Use:   "serve-http",
	Short: "Start the HTTP service",
	Run:   serveHttp,
}

func init() {
	serveHTTPCommand.Flags().StringP(flagBind, flagBindShort, flagBindDefault, flagBindDescr)
	serveHTTPCommand.Flags().IntP(flagPort, flagPortShort, flagPortDefault, flagPortDescr)

	_ = viper.BindPFlag(envFlagBind, serveHTTPCommand.Flags().Lookup(flagBind))
	_ = viper.BindPFlag(envFlagPort, serveHTTPCommand.Flags().Lookup(flagPort))

	rootCommand.AddCommand(serveHTTPCommand)
}

func serveHttp(cmd *cobra.Command, args []string) {
}
