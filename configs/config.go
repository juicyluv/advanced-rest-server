package configs

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// LoadConfigs loads .env file and config from given path.
// Returns an error if failed.
func LoadConfigs(configPath, configName string) error {
	err := loadViperConfig(configPath, configName)
	if err != nil {
		return err
	}

	return godotenv.Load()
}

// loadViperConfig tries to find config file and load it.
func loadViperConfig(configPath, configName string) error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)

	return viper.ReadInConfig()
}
