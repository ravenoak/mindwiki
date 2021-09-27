package cmd

import (
	"errors"
	"os"
	"os/signal"

	"github.com/ravenoak/mindwiki/app"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagWebUIBind        = "webui-bind"
	flagWebUIBindShort   = "b"
	flagWebUIBindDefault = "localhost"
	flagWebUIBindDescr   = "bind address for webui server"

	flagWebUIDaemon        = "webui-daemon"
	flagWebUIDaemonShort   = "z"
	flagWebUIDaemonDefault = false
	flagWebUIDaemonDescr   = "daemonize the webui server"

	flagWebUIPort        = "webui-port"
	flagWebUIPortShort   = "p"
	flagWebUIPortDefault = 1323
	flagWebUIPortDescr   = "port for webui server"

	envFlagWebUIBind   = "webui_bind"
	envFlagWebUIDaemon = "webui_daemon"
	envFlagWebUIPort   = "webui_port"
)

var serveWebUICommand = &cobra.Command{
	Use:   "serve-webui",
	Short: "Start the WebUI",
	Run:   webUI,
}

func init() {
	serveWebUICommand.Flags().StringP(flagWebUIBind, flagWebUIBindShort, flagWebUIBindDefault, flagWebUIBindDescr)
	serveWebUICommand.Flags().BoolP(flagWebUIDaemon, flagWebUIDaemonShort, flagWebUIDaemonDefault, flagWebUIDaemonDescr)
	serveWebUICommand.Flags().IntP(flagWebUIPort, flagWebUIPortShort, flagWebUIPortDefault, flagWebUIPortDescr)

	_ = viper.BindPFlag(envFlagWebUIBind, serveWebUICommand.Flags().Lookup(flagWebUIBind))
	_ = viper.BindPFlag(envFlagWebUIDaemon, serveWebUICommand.Flags().Lookup(flagWebUIDaemon))
	_ = viper.BindPFlag(envFlagWebUIPort, serveWebUICommand.Flags().Lookup(flagWebUIPort))

	rootCommand.AddCommand(serveWebUICommand)
}

func webUI(cmd *cobra.Command, args []string) {
	cfg := &app.Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatal().Err(err)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.DebugMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Warn().Interface("app.Config", cfg).Msg("")
	}

	mw := &app.App{
		Config: cfg,
	}

	go func() {
		if err := mw.StartWebUI(); err != nil {
			log.Fatal().Err(err)
		}
	}()

	if !cfg.WebUIDaemon {
		log.Info().Msg("not spawning to background")
		c := make(chan os.Signal, 1)
		// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
		// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
		signal.Notify(c, os.Interrupt)

		// Block until we receive our signal.
		<-c
		log.Info().Msg("shutting down")
		log.Fatal().Err(mw.StopWebUI())
	} else {
		log.Error().Err(errors.New("unsupported at this time, gonna shut down now"))
		log.Info().Msg("shutting down")
		log.Fatal().Err(mw.StopWebUI())
		os.Exit(1)
	}
	os.Exit(0)
}
