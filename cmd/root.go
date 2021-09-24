package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	envPrefix = "mindwiki"

	flagAppEnvDefault = "development"

	flagDebug        = "debug"
	flagDebugShort   = "d"
	flagDebugDefault = false
	flagDebugDescr   = "enable debug mode"

	flagStorage        = "storage"
	flagStorageShort   = "s"
	flagStorageDefault = "bbolt"
	flagStorageDescr   = "storage type"

	envFlagAppEnv  = "app_env"
	envFlagDebug   = "debug_mode"
	envFlagStorage = "storage"
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
	rootCommand.PersistentFlags().StringP(flagStorage, flagStorageShort, flagStorageDefault, flagStorageDescr)
}

func initConfig() {
	viper.SetEnvPrefix(envPrefix)

	viper.SetDefault(envFlagAppEnv, flagAppEnvDefault)
	viper.SetDefault(envFlagDebug, flagDebugDefault)

	_ = viper.BindPFlag(envFlagDebug, rootCommand.Flags().Lookup(flagDebug))
	_ = viper.BindPFlag(envFlagStorage, rootCommand.Flags().Lookup(flagStorage))

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/mindwiki/")
	viper.AddConfigPath("$HOME/.mindwiki/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	// TODO: Make this not panic (generate default? just run with default?) or more helpful/human message
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
