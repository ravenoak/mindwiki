package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/ravenoak/mindwiki/internal/app"
	"github.com/ravenoak/mindwiki/internal/config"
	//"github.com/ravenoak/mindwiki/internal/server"
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
	FLAG_BIND         = "bind"
	FLAG_BIND_SHORT   = "b"
	FLAG_BIND_DEFAULT = "0.0.0.0"
	FLAG_BIND_DESCR   = "bind address for web server"

	FLAG_PORT         = "port"
	FLAG_PORT_SHORT   = "p"
	FLAG_PORT_DEFAULT = 1323
	FLAG_PORT_DESCR   = "port for web server"

	FLAG_TLS_DISABLE         = "disable-tls"
	FLAG_TLS_DISABLE_DEFAULT = false
	FLAG_TLS_DISABLE_DESCR   = "disable TLS for web server"

	ENV_FLAG_BIND        = FLAG_BIND
	ENV_FLAG_PORT        = FLAG_PORT
	ENV_FLAG_TLS_DISABLE = "tls_disabled"
)

var serveHTTPCommand = &cobra.Command{
	Use:   "serve-http",
	Short: "Start the HTTP service",
	Run:   serveHttp,
}

func init() {
	serveHTTPCommand.Flags().StringP(FLAG_BIND, FLAG_BIND_SHORT, FLAG_BIND_DEFAULT, FLAG_BIND_DESCR)
	serveHTTPCommand.Flags().IntP(FLAG_PORT, FLAG_PORT_SHORT, FLAG_PORT_DEFAULT, FLAG_PORT_DESCR)
	serveHTTPCommand.Flags().Bool(FLAG_TLS_DISABLE, FLAG_TLS_DISABLE_DEFAULT, FLAG_TLS_DISABLE_DESCR)

	_ = viper.BindPFlag(ENV_FLAG_BIND, serveHTTPCommand.Flags().Lookup(FLAG_BIND))
	_ = viper.BindPFlag(ENV_FLAG_PORT, serveHTTPCommand.Flags().Lookup(FLAG_PORT))
	_ = viper.BindPFlag(ENV_FLAG_TLS_DISABLE, serveHTTPCommand.Flags().Lookup(FLAG_TLS_DISABLE))

	rootCommand.AddCommand(serveHTTPCommand)
}

func serveHttp(cmd *cobra.Command, args []string) {
	c := new(config.AppConfig)
	if err := viper.Unmarshal(c); err != nil {
		log.Fatal().Err(err)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if c.DebugMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Info().Str("config", fmt.Sprintf("%+v", c)).Msg("debug enabled")
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

func setupStorage(c config.StorageConfig) app.Storinator {
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

func startStorage(s app.Storinator) {
	log.Info().Msg("starting storage")
	err := s.Open()
	if err != nil {
		log.Fatal().Err(err)
	}
}

func stopStorage(s app.Storinator) {
	log.Info().Msg("stopping storage")
	if err := s.Close(); err != nil {
		log.Fatal().Err(err)
	}
}

func setupServer(c *config.AppConfig, s app.Storinator) app.HTTPServinator {
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
