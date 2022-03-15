package cmd

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ravenoak/mindwiki/app"
	storage2 "github.com/ravenoak/mindwiki/storage"
)

var storageCommand = &cobra.Command{
	Use:   "storage",
	Short: "Storage stuff",
}

var storageAddCommand = &cobra.Command{
	Use:   "add",
	Short: "Add to DB",
	Run:   storageAdd,
}

var storageReadCommand = &cobra.Command{
	Use:   "read",
	Short: "Read from DB",
	Run:   storageRead,
}

var storageDBStatsCommand = &cobra.Command{
	Use:   "db-stats",
	Short: "Database stats",
	Run:   storageStats,
}

func init() {
	storageCommand.AddCommand(storageAddCommand)
	storageCommand.AddCommand(storageReadCommand)
	storageCommand.AddCommand(storageDBStatsCommand)
	rootCommand.AddCommand(storageCommand)
}

func storageAdd(cmd *cobra.Command, args []string) {
	cfg := &app.Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatal().Err(err)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.DebugMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Warn().Interface("app.Config", cfg).Msg("")
	}

	log.Info().Str("cfg.StoragePath", cfg.StoragePath).Msg("")
	s, err := storage2.NewStorage(cfg.StoragePath)
	if err != nil {
		log.Fatal().Err(err)
	}

	if err = s.Open(); err != nil {
		log.Fatal().Err(err)
	}

	tx, err := s.NewTx(true)
	if err != nil {
		log.Fatal().Err(err)
	}

	b, err := tx.CreateBucketIfNotExists([]byte("testing"))
	if err != nil {
		log.Fatal().Err(err)
	}

	if err = b.Put([]byte("foo"), []byte("bar")); err != nil {
		log.Error().Err(err)
	}

	if err = tx.Commit(); err != nil {
		log.Error().Err(err)
	}

}

func storageRead(cmd *cobra.Command, args []string) {
	cfg := &app.Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatal().Err(err)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.DebugMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Warn().Interface("app.Config", cfg).Msg("")
	}

	s, err := storage2.NewStorage(cfg.StoragePath)
	if err != nil {
		log.Fatal().Err(err)
	}

	if err = s.Open(); err != nil {
		log.Fatal().Err(err)
	}

	tx, err := s.NewTx(false)
	if err != nil {
		log.Fatal().Err(err)
	}

	b := tx.Bucket([]byte("testing"))
	for _, i := range args {
		if _, err = fmt.Fprintf(cmd.OutOrStdout(), "%s: %+v\n", i, b.Get([]byte(i))); err != nil {
			log.Fatal().Err(err)
		}
	}

}

func storageStats(cmd *cobra.Command, args []string) {
	cfg := &app.Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatal().Err(err)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.DebugMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Warn().Interface("app.Config", cfg).Msg("")
	}

	s, err := storage2.NewStorage(cfg.StoragePath)
	if err != nil {
		log.Fatal().Err(err)
	}

	if err = s.Open(); err != nil {
		log.Fatal().Err(err)
	}

	if _, err = fmt.Fprintf(cmd.OutOrStdout(), "DB Stats: %s\n\n", s.DBStats()); err != nil {
		log.Fatal().Err(err)
	}
}
