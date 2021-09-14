package cmd

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/ravenoak/mindwiki/internal/config"
	"github.com/ravenoak/mindwiki/internal/server"
	"github.com/ravenoak/mindwiki/internal/storage"
	"github.com/ravenoak/mindwiki/internal/storage/adapters/bbolt"
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
	Run:   serve_http,
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

func serve_http(cmd *cobra.Command, args []string) {
	c := new(config.AppConfig)

	if err := viper.Unmarshal(c); err != nil {
		log.Fatal().Err(err)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if c.DebugMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Debug().Msg("Debug enabled")


	store_ := bbolt.NewBBoltAdapter(c.StorageConfig)
	store := storage.NewDepot(store_)

	log.Info().Msg("Initializing storage")
	err := store.Open()
	if err != nil {
		log.Fatal().Err(err)
	}

	s, err := server.New(c, store)
	if err != nil {
		log.Fatal().Err(err)
	}
	defer store.Close()

	log.Info().Msg("Starting server")
	go func() {
		if err := s.Start(); err != nil {
			log.Fatal().Err(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal().Err(err)
	}
}
