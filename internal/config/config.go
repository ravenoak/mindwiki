package config

import (
	"github.com/ravenoak/mindwiki/internal/validator"
)

type (
	AppConfig struct {
		AppEnv        string         `mapstructure:"app_env"`
		DebugMode     bool           `mapstructure:"debug_mode"`
		StorageConfig *StorageConfig `mapstructure:"storage_config"`
		WebConfig     *WebConfig     `mapstructure:"web_config"`
	}

	StorageConfig struct {
		Driver string `mapstructure:"driver"`
		Path   string `mapstructure:"path"`
	}

	WebConfig struct {
		Bind        string `mapstructure:"bind"`
		Port        uint16 `mapstructure:"port"`
		TLSDisabled bool   `mapstructure:"tls_disabled"`
	}
)

func (c *AppConfig) IsValid(errors validator.Error) {
	if c.AppEnv == "" {
		errors.Add("app environment not configured")
	}
	if c.StorageConfig == nil {
		errors.Add("storage not configured")
	}
	if c.WebConfig == nil {
		errors.Add("webserver not configured")
	}
}

func (c *StorageConfig) IsValid(errors validator.Error) {
	if c.Driver == "" {
		errors.Add("driver must not be blank")
	}
	if c.Path == "" {
		errors.Add("path must not be blank")
	}
}

func (c *WebConfig) IsValid(errors validator.Error) {
	if c.Port == 0 {
		errors.Add("port not configured")
	}
}
