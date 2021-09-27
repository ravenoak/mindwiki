package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	envPrefix = "mindwiki"

	flagDebug        = "debug"
	flagDebugShort   = "d"
	flagDebugDefault = false
	flagDebugDescr   = "enable debug mode"

	flagStorage      = "storage-path"
	flagStorageShort = "s"
	flagStorageDescr = "path to store all data"

	envFlagDebug   = "debug_mode"
	envFlagStorage = "storage_path"
)

var rootCommand = &cobra.Command{
	Use:   "mindwiki",
	Short: "A Secondary Brain in the shape of a Personal Wiki.",
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCommand.PersistentFlags().BoolP(flagDebug, flagDebugShort, flagDebugDefault, flagDebugDescr)
	rootCommand.PersistentFlags().StringP(flagStorage, flagStorageShort, "", flagStorageDescr)
}

func initConfig() {
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	viper.SetDefault(envFlagDebug, flagDebugDefault)

	_ = viper.BindPFlag(envFlagDebug, rootCommand.Flags().Lookup(flagDebug))
	_ = viper.BindPFlag(envFlagStorage, rootCommand.Flags().Lookup(flagStorage))

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/mindwiki/")
	viper.AddConfigPath("$HOME/.mindwiki/")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Info().Msg("config file not found")
		} else {
			// Config file was found but another error was produced
			log.Error().Err(err)
		}
	}
}
