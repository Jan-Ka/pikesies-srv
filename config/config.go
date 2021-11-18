package config

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"port"`

	GCPServiceAccountPath string `mapstructure:"gcp_service_account_path"`

	WAAppKeySecretKey string `mapstructure:"wa_app_key_secret_key"`
	WAAppSecret       string `mapstructure:"wa_app_key_secret"`
}

var lock = &sync.Mutex{}

type configManager struct {
	Config *Config
}

var configManagerInstance *configManager

func GetConfigManager() *configManager {
	if configManagerInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		cmLog := log.With().Str("package", "config").Logger()

		// cwd
		viper.AddConfigPath(".")

		// run path
		exPath, err := os.Executable()

		if err != nil {
			cmLog.Error().Msg("Unable to locate executable")
		}

		exePath := filepath.Dir(exPath)

		viper.AddConfigPath(exePath)

		// home path
		homePath, err := os.UserHomeDir()

		if err != nil {
			cmLog.Error().Msg("Unable to get user home dir")
		}

		viper.AddConfigPath(homePath)

		viper.SetConfigType("yaml")
		viper.SetConfigName(".pikesies-srv")

		configManagerInstance = &configManager{}

		cfg := &Config{
			Port: "8080",
		}

		err = viper.ReadInConfig()
		if err != nil {
			cmLog.Error().Msgf("Unable to read in config file due to %v", err)
		} else {
			if err = viper.Unmarshal(&cfg); err != nil {
				cmLog.Error().Msgf("Unable to unmarshal config file due to %v", err)
			}
		}

		configManagerInstance.Config = cfg
	}

	return configManagerInstance
}
