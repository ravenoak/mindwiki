package app

type Config struct {
	DebugMode   bool   `mapstructure:"debug_mode"`
	StoragePath string `mapstructure:"storage_path"`
	WebUIBind   string `mapstructure:"webui_bind"`
	WebUIDaemon bool   `mapstructure:"webui_daemon"`
	WebUIPort   uint16 `mapstructure:"webui_port"`
}
