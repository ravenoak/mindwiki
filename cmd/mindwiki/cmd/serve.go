package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/ravenoak/mindwiki/internal/app"
	"github.com/ravenoak/mindwiki/internal/config"
	"github.com/ravenoak/mindwiki/internal/validator"

	// "github.com/ravenoak/mindwiki/internal/server"
	server "github.com/ravenoak/mindwiki/internal/server/echo"
	"github.com/ravenoak/mindwiki/internal/storage"
	"github.com/ravenoak/mindwiki/internal/storage/adapters/bbolt"
	"github.com/ravenoak/mindwiki/internal/storage/adapters/sqlite"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	flagTLSDisable        = "disable-tls"
	flagTLSDisableDefault = false
	flagTLSDisableDescr   = "disable TLS for web server"

	envFlagBind       = flagBind
	envFlagPort       = flagPort
	envFlagTlsDisable = "tls_disabled"
)

var serveHTTPCommand = &cobra.Command{
	Use:   "serve-http",
	Short: "Start the HTTP service",
	Run:   serveHttp,
}

func init() {
	serveHTTPCommand.Flags().StringP(flagBind, flagBindShort, flagBindDefault, flagBindDescr)
	serveHTTPCommand.Flags().IntP(flagPort, flagPortShort, flagPortDefault, flagPortDescr)
	serveHTTPCommand.Flags().Bool(flagTLSDisable, flagTLSDisableDefault, flagTLSDisableDescr)

	_ = viper.BindPFlag(envFlagBind, serveHTTPCommand.Flags().Lookup(flagBind))
	_ = viper.BindPFlag(envFlagPort, serveHTTPCommand.Flags().Lookup(flagPort))
	_ = viper.BindPFlag(envFlagTlsDisable, serveHTTPCommand.Flags().Lookup(flagTLSDisable))

	rootCommand.AddCommand(serveHTTPCommand)
}

func serveHttp(cmd *cobra.Command, args []string) {
	c := new(config.AppConfig)
	if err := viper.Unmarshal(c); err != nil {
		log.Fatal().Err(err)
	}

	if err := validator.Validate(c); err != nil {
		log.Fatal().Err(err)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if c.DebugMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Warn().Str("config", fmt.Sprintf("%#v", c)).Msg("")
	}

	st := setupStorage(*c.StorageConfig)
	startStorage(st)
	s := setupServer(c, st)
	defer stopStorage(st)
	startServer(s)
	log.Debug().Msg("this is in that weird area I don't fully understand yet")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	log.Info().Msg("stopping server")
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal().Err(err)
	}
}

func setupStorage(c config.StorageConfig) app.Persistenator {
	log.Info().Msg("initializing storage")
	switch c.Driver {
	case "bbolt":
		return storage.NewDepot(bbolt.NewBBoltAdapter(&c))
	case "sqlite":
		return storage.NewDepot(sqlite.NewSQLiteAdapter(&c))
	case "gorp-sqlite":
		return storage.NewORP(&c)
	default:
		log.Fatal().Err(app.InvalidStorageTypeError(c.Driver))
	}
	return nil
}

func startStorage(s app.Persistenator) {
	log.Info().Msg("starting storage")
	err := s.Open()
	if err != nil {
		log.Fatal().Err(err)
	}
}

func stopStorage(s app.Persistenator) {
	log.Info().Msg("stopping storage")
	if err := s.Close(); err != nil {
		log.Fatal().Err(err)
	}
}

func setupServer(c *config.AppConfig, s app.Persistenator) app.HTTPServinator {
	log.Debug().Msg("initializing server")
	h, err := server.HTTPServer(c, s)
	if err != nil {
		log.Fatal().Err(err)
	}
	return h
}

func startServer(s app.HTTPServinator) {
	log.Info().Msg("starting server")
	go func() {
		if err := s.Start(); err != nil {
			log.Fatal().Err(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
