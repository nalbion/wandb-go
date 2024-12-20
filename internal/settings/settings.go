package settings

import "os"

// adapted from https://github.com/wandb/wandb-js/blob/main/src/sdk/settings.ts
// TODO: convert to struct & flesh out as needed - https://github.com/wandb/wandb/blob/main/wandb/old/settings.py#L13
type Settings struct {
	APIKey  string
	BaseURL string
	Offline bool
	Silent  bool
}

func (s *Settings) GetFileStreamMaxBytes() int {
	return 1024 * 1024
}

type Config struct {
	Debug  bool
	Env    string
	Mode   string // online, offline, disabled
	APIKey string
	APIURL string
	AppURL string
	// WinHome string HOME || HOMEDRIVE\HOMEPATH || USERPROFILE
}

func NewConfig() *Config {
	env := os.Getenv("WANDB_ENV")
	if env == "" {
		env = "production"
	}

	mode := os.Getenv("WANDB_MODE")
	if mode == "" {
		mode = "online"
	}

	return &Config{
		Debug:  os.Getenv("WANDB_DEBUG") == "true",
		Env:    env,
		Mode:   mode,
		APIKey: os.Getenv("WANDB_API_KEY"),
		APIURL: os.Getenv("WANDB_BASE_URL"),
		AppURL: os.Getenv("WANDB_APP_URL"),
	}
}

func DefaultSettings(cfg *Config) *Settings {
	baseUrl := cfg.APIURL
	if baseUrl == "" {
		baseUrl = "https://api.wandb.ai"
	}

	return &Settings{
		APIKey:  cfg.APIKey,
		BaseURL: baseUrl,
		Offline: cfg.Mode == "offline",
	}
}

func SettingsWithOverrides(cfg *Config, overrides *Settings) *Settings {
	settings := DefaultSettings(cfg)

	if overrides.APIKey != "" {
		settings.APIKey = overrides.APIKey
	}

	if overrides.BaseURL != "" {
		settings.BaseURL = overrides.BaseURL
	}

	if overrides.Offline {
		settings.Offline = overrides.Offline
	}

	return settings
}
